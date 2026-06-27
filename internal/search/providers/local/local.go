package local

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"path"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/search/query"
	"github.com/tanamoe/urano/internal/config"
	"github.com/tanamoe/urano/internal/models"
	"github.com/tanamoe/urano/internal/search"
)

const (
	RegistryIndexName = "registry.bleve"
)

type bleveProvider struct {
	cfg  config.AppConfig
	repo *models.Queries

	registryIndex bleve.Index
}

func NewBleveProvider(cfg config.AppConfig, repo *models.Queries) (search.Search, error) {
	registryPath := path.Join(cfg.StatePath, RegistryIndexName)

	registryIndex, err := bleve.New(registryPath, bleve.NewIndexMapping())
	if errors.Is(err, bleve.ErrorIndexPathExists) {
		registryIndex, err = bleve.Open(registryPath)
	}
	if err != nil {
		return nil, err
	}

	return &bleveProvider{
		cfg:           cfg,
		repo:          repo,
		registryIndex: registryIndex,
	}, nil
}

func (p *bleveProvider) Stop(ctx context.Context) error {
	if err := p.registryIndex.Close(); err != nil {
		return err
	}

	return nil
}

func (p *bleveProvider) IndexAllRegistry(ctx context.Context) error {
	var count int32

	for offset := int32(0); ; offset += p.cfg.IndexBatchSize {
		registries, err := p.repo.ListRegistry(ctx, models.ListRegistryParams{
			Limit:  p.cfg.IndexBatchSize,
			Offset: offset,
		})
		if err != nil {
			return err
		}
		count += int32(len(registries))

		batch := p.registryIndex.NewBatch()

		for _, registry := range registries {
			if err := batch.Index(registry.ID.String(), registry); err != nil {
				return err
			}
		}

		slog.DebugContext(ctx, fmt.Sprintf("indexing offset %d with batch size %d", offset, p.cfg.IndexBatchSize))
		if err := p.registryIndex.Batch(batch); err != nil {
			return err
		}

		if int32(len(registries)) < p.cfg.IndexBatchSize {
			break
		}
	}

	slog.DebugContext(ctx, "completed indexing registries", "count", count)

	return nil
}

func (p *bleveProvider) IndexRegistry(ctx context.Context, registry *models.Registry) error {
	if registry == nil {
		return errors.New("registry is empty")
	}

	return p.registryIndex.Index(registry.ID.String(), registry)
}

func (p *bleveProvider) SearchRegistry(ctx context.Context, request search.SearchRegistryRequest) ([]string, error) {
	query, err := buildSearchRegistryQuery(request)
	if err != nil {
		return nil, err
	}

	search := bleve.NewSearchRequest(query)
	result, err := p.registryIndex.Search(search)
	if err != nil {
		return nil, err
	}

	hits := make([]string, 0, result.Hits.Len())
	for _, hit := range result.Hits {
		hits = append(hits, hit.ID)
	}

	return hits, nil
}

func buildSearchRegistryQuery(request search.SearchRegistryRequest) (query.Query, error) {
	queries := []query.Query{}

	if request.RegistrationID != nil {
		q := query.NewTermQuery(*request.RegistrationID)
		q.SetField("registrationID")
		queries = append(queries, q)
	}

	if request.ISBN != nil {
		q := query.NewTermQuery(*request.ISBN)
		q.SetField("isbn")
		queries = append(queries, q)
	}

	if request.Title != nil {
		q := query.NewMatchQuery(*request.Title)
		q.SetField("title")
		queries = append(queries, q)
	}

	if request.Author != nil {
		q := query.NewMatchQuery(*request.Author)
		q.SetField("author")
		queries = append(queries, q)
	}

	if request.Translator != nil {
		q := query.NewMatchQuery(*request.Translator)
		q.SetField("translator")
		queries = append(queries, q)
	}

	if request.PrintAmount != nil {
		q := NumericRangeQuery(*request.PrintAmount)
		q.SetField("printAmount")
		queries = append(queries, q)
	}

	if request.SelfPublish != nil {
		q := query.NewBoolFieldQuery(*request.SelfPublish)
		q.SetField("selfPublish")
		queries = append(queries, q)
	}

	if request.Partner != nil {
		q := query.NewMatchQuery(*request.Partner)
		q.SetField("partner")
		queries = append(queries, q)
	}

	return query.NewConjunctionQuery(queries), nil
}
