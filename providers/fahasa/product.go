package fahasa

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
)

const (
	productPath = "/products/id"
)

func (c *client) Product(ctx context.Context, productID int64) (*Product, error) {
	id := strconv.FormatInt(productID, 10)
	path, err := url.JoinPath(c.domain, productPath, id)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", defaultUserAgent)

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			slog.Warn("cannot close response body", "error", err)
		}
	}()

	if response.StatusCode != http.StatusOK {
		slog.Error("downstream returns invalid status code", "code", response.StatusCode)
		return nil, errors.New("downstream failed")
	}

	var product Product
	if err := json.NewDecoder(response.Body).Decode(&product); err != nil {
		return nil, err
	}

	return &product, nil
}
