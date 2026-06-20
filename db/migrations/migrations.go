package migrations

import (
	"context"
	"database/sql"
	"embed"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed *.sql
var migrations embed.FS

type Migration interface {
	Run(ctx context.Context) error
}

type migration struct {
	db *sql.DB
}

func NewMigration(db *sql.DB) Migration {
	return &migration{db}
}

func NewMigrationFromPool(dbPool *pgxpool.Pool) Migration {
	db := stdlib.OpenDBFromPool(dbPool)

	return &migration{db: db}
}

// runs migration if present
func (m *migration) Run(ctx context.Context) error {
	goose.SetBaseFS(migrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.UpContext(ctx, m.db, "."); err != nil {
		return err
	}

	return m.db.Close()
}
