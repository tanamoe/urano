package main

import (
	"net/http"

	"buf.build/gen/go/tanamoe/urano/connectrpc/go/urano/api/v1beta1/apiv1beta1connect"
	connectcors "connectrpc.com/cors"
	"github.com/rs/cors"
	"github.com/tanamoe/urano/internal/service/v1beta1"
)

func main() {
	aggregate := v1beta1.NewAggregateServer()

	mux := http.NewServeMux()
	path, handler := apiv1beta1connect.NewAggregateServiceHandler(
		aggregate,
	)
	handler = withCORS(handler)
	mux.Handle(path, handler)
	p := new(http.Protocols)
	p.SetHTTP1(true)
	p.SetUnencryptedHTTP2(true)
	s := http.Server{
		Addr:      "0.0.0.0:8080",
		Handler:   mux,
		Protocols: p,
	}
	s.ListenAndServe()
}

// withCORS adds CORS support to a Connect HTTP handler.
func withCORS(connectHandler http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // replace with your domain
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: connectcors.AllowedHeaders(),
		ExposedHeaders: connectcors.ExposedHeaders(),
		MaxAge:         7200, // 2 hours in seconds
	})
	return c.Handler(connectHandler)
}
