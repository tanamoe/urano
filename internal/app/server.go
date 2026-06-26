package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"buf.build/gen/go/tanamoe/urano/connectrpc/go/urano/api/v1beta1/apiv1beta1connect"
	"connectrpc.com/connect"
	connectcors "connectrpc.com/cors"
	"connectrpc.com/otelconnect"
	"github.com/rs/cors"
	"github.com/tanamoe/urano/internal/service/v1beta1"
)

func (a *app) initServer(_ context.Context) error {
	interceptors, err := a.setupInterceptors()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	path, handler := apiv1beta1connect.NewAggregateServiceHandler(
		v1beta1.NewAggregateServer(a.repo, a.cfg.Fahasa.SearchToken),
		connect.WithInterceptors(interceptors...),
	)
	handler = withCORS(a.cfg.App.AllowedOrigins, handler)
	mux.Handle(path, handler)

	path, handler = apiv1beta1connect.NewRegistryServiceHandler(
		v1beta1.NewRegistryServer(a.repo),
		connect.WithInterceptors(interceptors...),
	)
	handler = withCORS(a.cfg.App.AllowedOrigins, handler)
	mux.Handle(path, handler)

	p := new(http.Protocols)
	p.SetHTTP1(true)
	p.SetUnencryptedHTTP2(true)

	a.server = &http.Server{
		Addr:      a.cfg.App.ListenAddress,
		Handler:   mux,
		Protocols: p,
	}

	a.onShutdown.BindFunc(func(ctx context.Context, _ any) error {
		return a.server.Shutdown(ctx)
	})

	return nil
}

func (a *app) setupInterceptors() ([]connect.Interceptor, error) {
	otelInterceptor, err := otelconnect.NewInterceptor()
	if err != nil {
		return nil, err
	}

	return []connect.Interceptor{
		otelInterceptor,
	}, nil
}

// withCORS wraps a Connect HTTP handler with CORS middleware.
func withCORS(allowedOrigins []string, h http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: allowedOrigins,
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: connectcors.AllowedHeaders(),
		ExposedHeaders: connectcors.ExposedHeaders(),
		MaxAge:         7200, // 2 hours in seconds
	})
	return c.Handler(h)
}

func (a *app) Serve(ctx context.Context) error {
	slog.InfoContext(ctx, "server listening", "addr", a.cfg.App.ListenAddress)

	if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("serve err: %w", err)
	}

	return nil
}
