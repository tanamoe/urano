package ppdvntest

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
)

//go:embed *.html
var html embed.FS

const (
	RegistriesPath = "/web/guest/ke-hoach-xuat-ban"

	QueryFrom        = "bat_dau"
	QueryTo          = "ket_thuc"
	QuerySearchQuery = "query"
	QueryPage        = "p"
)

func NewServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc(RegistriesPath, registriesHandler)
	return httptest.NewServer(mux)
}

func registriesHandler(w http.ResponseWriter, req *http.Request) {
	queries := req.URL.Query()

	query := queries.Get(QuerySearchQuery)
	page := queries.Get(QueryPage)

	var pageInt int
	var err error
	if page != "" {
		pageInt, err = strconv.Atoi(page)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	if strings.Contains(query, "CLAMP") {
		serveHTML(w, "single-page-registries.html")
		return
	}

	if strings.Contains(query, "Cửu Long") {
		serveHTML(w, "kowloon.html")
		return
	}

	if strings.Contains(query, "Card Captor Sakura") && pageInt >= 1 && pageInt <= 10 {
		serveHTML(w, fmt.Sprintf("ccs-%d.html", pageInt))
		return
	}

	serveHTML(w, "no-registries.html")
}

func serveHTML(w http.ResponseWriter, path string) {
	file, err := html.Open(path)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(w, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
