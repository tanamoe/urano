package fahasa

import (
	"context"
	"crypto/tls"
	"net/http"
)

const (
	defaultRestDomain = "https://rest.fahasa.com"
	defaultBaseDomain = "https://www.fahasa.com"
	defaultUserAgent  = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36"
)

type Client interface {
	Product(ctx context.Context, productID int64) (*Product, error)
	ListByCategory(ctx context.Context, params ListByCategoryParams) (*CategoryProducts, error)
	Search(ctx context.Context, query string) (*SearchResponse, error)
}

type ListByCategoryParams struct {
	CategoryID int64

	Page     int32
	PageSize int32
}

type client struct {
	restDomain  string
	baseDomain  string
	searchToken string

	httpClient *http.Client
}

type clientOptions = func(client *client)

func NewClient(options ...clientOptions) Client {
	client := &client{
		restDomain: defaultRestDomain,
		baseDomain: defaultBaseDomain,
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

func WithRestDomain(domain string) clientOptions {
	return func(client *client) {
		client.restDomain = domain
	}
}

func WithBaseDomain(domain string) clientOptions {
	return func(client *client) {
		client.baseDomain = domain
	}
}

// search token starts with search-
func WithSearchToken(searchToken string) clientOptions {
	return func(client *client) {
		client.searchToken = searchToken
	}
}
