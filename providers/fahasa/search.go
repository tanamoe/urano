package fahasa

import (
	"context"
	"encoding/json/jsontext"
	"encoding/json/v2"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

const (
	searchPath = "/api/elsearch/api/as/v1/engines/fhs-production-v2/search.json"
)

func (c *client) Search(ctx context.Context, query string) (*SearchResponse, error) {
	r := strings.NewReader(fmt.Sprintf(`{"query": "%s", "sort": { "created_at": "desc" }, "page": { "size": 48, "current": 1 }}`, query))

	path, err := url.JoinPath(c.baseDomain, searchPath)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, path, r)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", defaultUserAgent)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.searchToken))
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
	decoder := jsontext.NewDecoder(response.Body)
	if err := json.UnmarshalDecode(decoder, &search); err != nil {
		return nil, err
	}

	return &search, nil
}
