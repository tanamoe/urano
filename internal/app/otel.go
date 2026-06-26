package app

import (
	"context"
	"log/slog"

	"github.com/tanamoe/urano/pkg/otel"
)

// setupOTelSDK bootstraps the OpenTelemetry pipeline.
// If it does not return an error, make sure to call shutdown for proper cleanup.
func (a *app) initOtel(ctx context.Context) error {
	if err := otel.Init(ctx); err != nil {
		return err
	}
	a.OnShutdown().BindFunc(func(ctx context.Context, _ any) error {
		slog.InfoContext(ctx, "shutting down opentelemetry")
		return otel.Shutdown(ctx)
	})

	return nil
}
