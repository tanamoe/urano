package app

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tanamoe/urano/db/migrations"
	"github.com/tanamoe/urano/internal/config"
	"github.com/tanamoe/urano/internal/hooks"
	"github.com/tanamoe/urano/internal/models"
)

type App interface {
	Serve(ctx context.Context) error
	Stop(ctx context.Context) error

	OnShutdown() *hooks.Hook[any]
}

type app struct {
	cfg *config.Config

	db *pgxpool.Pool

	repo   *models.Queries
	server *http.Server

	onShutdown *hooks.Hook[any]
}

func New(ctx context.Context, cfg *config.Config) (App, error) {
	a := &app{
		cfg: cfg,
	}

	// init hooks
	a.onShutdown = &hooks.Hook[any]{}

	if err := a.initOtel(ctx); err != nil {
		return nil, err
	}

	if err := a.initDB(ctx); err != nil {
		return nil, err
	}

	// TODO: might refactor run migrations to be in serve
	if err := a.runMigrations(ctx); err != nil {
		return nil, err
	}

	if err := a.initServer(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

// Stop performs a graceful shutdown of all resources.
func (a *app) Stop(ctx context.Context) error {
	return a.onShutdown.Trigger(ctx, nil)
}

func (a *app) OnShutdown() *hooks.Hook[any] {
	return a.onShutdown
}

func (a *app) initDB(ctx context.Context) error {
	db, err := pgxpool.New(ctx, a.cfg.Database.Conn)
	if err != nil {
		return err
	}

	a.OnShutdown().BindFunc(func(ctx context.Context, a any) error {
		db.Close()
		return nil
	})

	a.db = db

	repo := models.New(db)
	a.repo = repo

	return nil
}

func (a *app) runMigrations(ctx context.Context) error {
	migration := migrations.NewMigrationFromPool(a.db)
	return migration.Run(ctx)
}
