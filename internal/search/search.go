package search

import (
	"context"

	"github.com/tanamoe/urano/internal/models"
	"golang.org/x/exp/constraints"
)

type Numeric interface {
	constraints.Float | constraints.Integer
}

type Search interface {
	Stop(context.Context) error

	SearchRegistry(ctx context.Context, query SearchRegistryRequest) ([]string, error)

	IndexRegistry(ctx context.Context, registry *models.Registry) error
	IndexAllRegistry(ctx context.Context) error
}

type QueryKind string

const (
	AndQuery QueryKind = "and"
	OrQuery  QueryKind = "or"
)

type QueryGroup[T comparable] struct {
	Kind   QueryKind `json:"kind"`
	Values []T       `json:"values"`
}

type QueryNumericRange[T Numeric] struct {
	Min *T `json:"min"`
	Max *T `json:"max"`
}
