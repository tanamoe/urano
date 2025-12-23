package fahasa

import (
	"context"
	"crypto/tls"
	"net/http"
)

const (
	defaultDomain    = "https://rest.fahasa.com"
	defaultUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36"
)

type Client interface {
	Product(ctx context.Context, productID int64) (*Product, error)
	ListByCategory(ctx context.Context, params ListByCategoryParams) (*CategoryProducts, error)
}

type ListByCategoryParams struct {
	CategoryID int64

	Page     int32
	PageSize int32
}

type client struct {
	domain string

	httpClient *http.Client
}

type clientOptions = func(client *client)

func NewClient(options ...clientOptions) Client {
	client := &client{
		domain: defaultDomain,
		httpClient: &http.Client{
			Transport: &http.Transport{
				TLSNextProto: map[string]func(string, *tls.Conn) http.RoundTripper{},
			},
		},
	}

	for _, option := range options {
		option(client)
	}

	return client
}

func WithDomain(domain string) clientOptions {
	return func(client *client) {
		client.domain = domain
	}
}
