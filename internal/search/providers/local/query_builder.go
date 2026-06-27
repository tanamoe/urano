package local

import (
	"fmt"

	"github.com/blevesearch/bleve/v2/search/query"
	"github.com/tanamoe/urano/internal/search"
)

func QueryGroupValues[T comparable](qg search.QueryGroup[T]) (vals []string) {
	for _, val := range qg.Values {
		vals = append(vals, fmt.Sprintf("%v", val))
	}
	return vals
}

func GroupQuery(kind search.QueryKind, queries []query.Query) query.Query {
	switch kind {
	case search.AndQuery:
		return query.NewConjunctionQuery(queries)
	case search.OrQuery:
		return query.NewDisjunctionQuery(queries)
	default:
		return query.NewMatchNoneQuery()
	}
}

func NumericRangeQuery[T search.Numeric](rq search.QueryNumericRange[T]) *query.NumericRangeQuery {
	var queryMin, queryMax *float64

	if rq.Min != nil {
		queryMin = new(float64(*rq.Min))
	}

	if rq.Max != nil {
		queryMax = new(float64(*rq.Max))
	}

	return query.NewNumericRangeQuery(queryMin, queryMax)
}
