package ppdvn

import (
	"context"
	"net/http"
	"time"
)

type Client interface {
	List(ctx context.Context, params ListParams) ([]Registry, error)
	GetLastPage(ctx context.Context, params ListParams) (int, error)
}

type ListParams struct {
	Query              string
	Page               int
	StartDate, EndDate time.Time
}

type client struct {
	httpClient *http.Client
}

func NewClient() Client {
	client := &client{
		httpClient: http.DefaultClient,
	}

	return client
}
