package providers

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"path"

	"github.com/blevesearch/bleve/v2"
	"github.com/tanamoe/urano/internal/config"
	"github.com/tanamoe/urano/internal/models"
	"github.com/tanamoe/urano/internal/search"
)

const (
	RegistryIndexName = "registry.bleve"
)

type bleveProvider struct {
	cfg           config.AppConfig
	repo          *models.Queries
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

func (p *bleveProvider) IndexAll(ctx context.Context) error {
	for offset := int32(0); ; offset += p.cfg.IndexBatchSize {
		registries, err := p.repo.ListRegistry(ctx, models.ListRegistryParams{
			Limit:  p.cfg.IndexBatchSize,
			Offset: offset,
		})
		if err != nil {
			return err
		}

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

	return nil
}

func (p *bleveProvider) SearchRegistry(ctx context.Context, q string) ([]models.Registry, error) {
	req := bleve.NewMatchQuery(q)
	search := bleve.NewSearchRequest(req)
	result, err := p.registryIndex.Search(search)
	if err != nil {
		return nil, err
	}

	slog.InfoContext(ctx, "got results", "results", result)

	return nil, nil
}
