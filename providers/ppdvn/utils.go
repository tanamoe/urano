package ppdvn

import (
	"strings"

	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
)

func getAttr(node *html.Node, key string) *html.Attribute {
	for _, a := range node.Attr {
		if a.Key == key {
			return &a
		}
	}

	return nil
}

func query(n *html.Node, query string) *html.Node {
	sel, err := cascadia.Parse(query)
	if err != nil {
		return &html.Node{}
	}
	return cascadia.Query(n, sel)
}

func text(n *html.Node) string {
	switch n.Type {
	case html.TextNode:
		return strings.Trim(n.Data, " ")
	default:
		var builder strings.Builder
		for n := range n.ChildNodes() {
			builder.WriteString(text(n))
		}
		return builder.String()
	}
}

func queryAll(n *html.Node, query string) []*html.Node {
	sel, err := cascadia.Parse(query)
	if err != nil {
		return []*html.Node{}
	}
	return cascadia.QueryAll(n, sel)
}
