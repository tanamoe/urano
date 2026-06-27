package search

import (
	"context"

	"github.com/tanamoe/urano/internal/models"
)

type Search interface {
	Stop(context.Context) error

	SearchRegistry(ctx context.Context, q string) ([]models.Registry, error)

	IndexAll(ctx context.Context) error
}
