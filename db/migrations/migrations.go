package migrations

import (
	"context"
	"database/sql"
	"embed"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed *.sql
var migrations embed.FS

type Migration struct {
	db *sql.DB
}

func NewMigrationFromPool(dbPool *pgxpool.Pool) *Migration {
	db := stdlib.OpenDBFromPool(dbPool)

	return &Migration{
		db: db,
	}
}

func (m *Migration) Close() error {
	return m.db.Close()
}

// runs migration if present
func (m *Migration) RunMigrations(ctx context.Context) error {
	goose.SetBaseFS(migrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	slog.InfoContext(ctx, "Running migrations...")

	if err := goose.UpContext(ctx, m.db, "."); err != nil {
		return err
	}

	slog.InfoContext(ctx, "Migration completed!")

	return nil
}
