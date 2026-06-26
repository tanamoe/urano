package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/pflag"
	"github.com/tanamoe/urano/internal/config"
	"github.com/tanamoe/urano/internal/models"
	"github.com/tanamoe/urano/providers/ppdvn"
)

var (
	from, to time.Time

	defaultFrom = time.Date(2016, time.January, 1, 0, 0, 0, 0, time.Local)
	defaultTo   = time.Now()
)

func main() {
	pflag.TimeVar(&from, "from", defaultFrom, []string{time.DateOnly}, "Date to filter registration from")
	pflag.TimeVar(&to, "to", defaultTo, []string{time.DateOnly}, "Date to filter registration to")
	pflag.Parse()

	ctx := context.Background()

	if err := run(ctx); err != nil {
		slog.ErrorContext(ctx, "uncoverable error occurred", "error", err.Error())
	}
}

func run(ctx context.Context) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	db, err := pgxpool.New(ctx, cfg.Database.Conn)
	if err != nil {
		return err
	}

	repo := models.New(db)

	client, err := ppdvn.NewClient()
	if err != nil {
		return err
	}

	for d := from; !d.After(to); d = d.AddDate(0, 0, 1) {
		last, err := client.GetLastPage(ctx, ppdvn.ListParams{
			Page:      new(1),
			StartDate: &d,
			EndDate:   &d,
		})
		if err != nil {
			slog.ErrorContext(ctx, "cannot retrieve last page", "error", err)
		}

		for i := 1; i <= last+1; i++ { // offset with 1
			results, err := client.List(ctx, ppdvn.ListParams{
				Page:      new(i + 1),
				StartDate: &d,
				EndDate:   &d,
			})
			if err != nil {
				slog.ErrorContext(ctx, "cannot retrieve registrations", "error", err, "date", d, "page", i)
			}

			slog.InfoContext(ctx, "successfully retrieve registrations", "date", d, "page", i, "count", len(results))

			for _, registry := range results {
				if _, err := repo.CreateRegistry(ctx, models.CreateRegistryParams{
					RegistrationID: registry.RegistrationID,
					Isbn: pgtype.Text{
						String: registry.ISBN,
						Valid:  true,
					},
					Title: registry.Title,
					Author: pgtype.Text{
						String: registry.Author,
						Valid:  true,
					},
					Translator: pgtype.Text{
						String: registry.Translator,
						Valid:  true,
					},
					PrintAmount: pgtype.Int4{
						Int32: int32(registry.PrintAmount),
						Valid: true,
					},
					SelfPublish: pgtype.Bool{
						Bool:  registry.SelfPublished,
						Valid: true,
					},
					Partner: pgtype.Text{
						String: registry.Partner,
						Valid:  true,
					},
					RegistrationDate: pgtype.Date{
						Time:  d,
						Valid: true,
					},
				}); err != nil {
					slog.Error(err.Error())
				}
			}
		}
	}

	return nil
}
