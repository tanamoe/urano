package ppdvn

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

const (
	DefaultHost   = "ppdvn.gov.vn"
	DefaultScheme = "https"
)

var (
	DefaultURL = url.URL{
		Host:   DefaultHost,
		Scheme: DefaultScheme,
	}
)

type Client interface {
	List(ctx context.Context, params ListParams) ([]Registry, error)
	GetLastPage(ctx context.Context, params ListParams) (int, error)
}

type ListParams struct {
	Query              *string
	Page               *int
	StartDate, EndDate *time.Time
}

type client struct {
	baseURL    url.URL
	httpClient *http.Client
}

type clientOptions = func(client *client) error

func NewClient(options ...clientOptions) (Client, error) {
	client := &client{
		baseURL:    DefaultURL,
		httpClient: http.DefaultClient,
	}

	for _, option := range options {
		if err := option(client); err != nil {
			return nil, err
		}
	}

	return client, nil
}

func WithHTTPClient(c *http.Client) clientOptions {
	return func(client *client) error {
		client.httpClient = c
		return nil
	}
}

func WithDomain(domain string) clientOptions {
	return func(client *client) error {
		baseURL, err := url.Parse(domain)
		if err != nil {
			return err
		}

		client.baseURL = *baseURL
		return nil
	}
}
