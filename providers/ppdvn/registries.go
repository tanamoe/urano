package ppdvn

import (
	"context"
	"log/slog"
	"net/http"
	"net/url"
	"regexp"
	"strconv"

	"golang.org/x/net/html"
)

const (
	Host           = "ppdvn.gov.vn"
	RegistriesPath = "/web/guest/ke-hoach-xuat-ban"

	dateFormat = "02/01/2006"

	noDataValue = "Không tìm thấy dữ liệu"
)

func (c *client) List(ctx context.Context, params ListParams) ([]Registry, error) {
	q := url.Values{}
	q.Add("p", strconv.Itoa(params.Page))
	q.Add("query", params.Query)
	q.Add("bat_dau", params.StartDate.Format(dateFormat))
	q.Add("ket_thuc", params.EndDate.Format(dateFormat))

	url := url.URL{
		Scheme:   "https",
		Host:     Host,
		Path:     RegistriesPath,
		RawQuery: q.Encode(),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //nolint:errcheck

	slog.DebugContext(ctx, "got response", "url", url.String())

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	node := query(doc, "div#list_data_return tbody")
	if node == nil {
		return []Registry{}, nil
	}

	return parseTable(node)
}

func parseTable(node *html.Node) ([]Registry, error) {
	nodes := queryAll(node, "tr")

	registries := make([]Registry, 0, len(nodes))
	for _, n := range nodes {
		if text(n) == noDataValue {
			return registries, nil
		}

		printAmount := text(query(n, "td:nth-child(6)"))

		i, err := strconv.Atoi(printAmount)
		if err != nil {
			return nil, err
		}

		registries = append(registries, Registry{
			ISBN:           text(query(n, "td:nth-child(2)")),
			Title:          text(query(n, "td:nth-child(3)")),
			Author:         text(query(n, "td:nth-child(4)")),
			Translator:     text(query(n, "td:nth-child(5)")),
			PrintAmount:    i,
			SelfPublished:  text(query(n, "td:nth-child(7)")) == "x",
			Partner:        text(query(n, "td:nth-child(8)")),
			RegistrationID: text(query(n, "td:nth-child(9)")),
		})
	}

	return registries, nil
}

func (c *client) GetLastPage(ctx context.Context, params ListParams) (int, error) {
	q := url.Values{}
	q.Add("p", strconv.Itoa(params.Page))
	q.Add("query", params.Query)
	q.Add("bat_dau", params.StartDate.Format(dateFormat))
	q.Add("ket_thuc", params.EndDate.Format(dateFormat))

	url := url.URL{
		Scheme:   "https",
		Host:     Host,
		Path:     RegistriesPath,
		RawQuery: q.Encode(),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return 0, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close() //nolint:errcheck

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return 0, err
	}

	node := query(doc, ".pagination ul li:last-child a[href]")
	if node == nil {
		return 0, nil
	}

	href := getAttr(node, "href")
	if href == nil {
		return 0, nil
	}

	r := regexp.MustCompile(`\d+$`)
	match := r.FindString(href.Val)
	page, err := strconv.Atoi(match)
	if err != nil {
		return 0, err
	}

	return page, nil
}
