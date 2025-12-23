package fahasa

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
)

const (
	categoryPath = "/category2"
)

func (c *client) ListByCategory(ctx context.Context, params ListByCategoryParams) (*CategoryProducts, error) {
	path, err := url.JoinPath(c.domain, categoryPath)
	if err != nil {
		return nil, err
	}

	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	queries := url.Query()

	queries.Add("catId", strconv.FormatInt(params.CategoryID, 10))
	queries.Add("sortBy", "created_at")
	queries.Add("pageSize", strconv.FormatInt(int64(params.PageSize), 10))
	queries.Add("page", strconv.FormatInt(int64(params.Page), 10))

	url.RawQuery = queries.Encode()

	fmt.Println(url.String())

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
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

	var products CategoryProducts
	if err := json.NewDecoder(response.Body).Decode(&products); err != nil {
		return nil, err
	}

	return &products, nil
}
