package fahasa

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"strings"
)

const (
	searchPath = "https://www.fahasa.com/api/elsearch/api/as/v1/engines/fhs-production-v2/search.json"
)

func (c *client) Search(ctx context.Context, query string) (*SearchResponse, error) {
	r := strings.NewReader(fmt.Sprintf(`{"query": "%s", "sort": { "created_at": "desc" }, "page": { "size": 48, "current": 1 }}`, query))

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, searchPath, r)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", defaultUserAgent)
	request.Header.Set("Authorization", "Bearer search-zs5622edvxg9bt9nb5y9m2oo")
	request.Header.Set("Content-Type", "application/json")

	req, _ := httputil.DumpRequestOut(request, true)
	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			slog.Warn("cannot close response body", "error", err)
		}
	}()
	res, _ := httputil.DumpResponse(response, true)

	if response.StatusCode != http.StatusOK {
		slog.Error("downstream returns invalid status code", "code", response.StatusCode, "request", req, "response", res)
		return nil, errors.New("downstream failed")
	}

	var search SearchResponse
	if err := json.NewDecoder(response.Body).Decode(&search); err != nil {
		return nil, err
	}

	return &search, nil
}
