package fahasa

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
)

const (
	defaultDomain string = "https://rest.fahasa.com"
	productURL    string = "/products/id/"
)

type Client interface {
	Product(ctx context.Context, productID int64) (*Product, error)
}

type client struct {
	domain string

	httpClient *http.Client
}

type clientOptions = func(client *client)

func NewClient(options ...clientOptions) Client {
	client := &client{
		domain:     defaultDomain,
		httpClient: http.DefaultClient,
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

func (c *client) Product(ctx context.Context, productID int64) (*Product, error) {
	url, err := url.JoinPath(c.domain, productURL, strconv.FormatInt(productID, 10))
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			slog.Warn("cannot close response body", "error", err)
		}
	}()

	var product Product
	if err := json.NewDecoder(response.Body).Decode(&product); err != nil {
		return nil, err
	}

	return &product, nil
}
