package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/tanamoe/urano/internal/app"
	"github.com/tanamoe/urano/internal/config"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cfg, err := config.Load()
	if err != nil {
		log.Panicln(err)
	}

	a, err := app.New(ctx, cfg)
	if err != nil {
		log.Panicln(err)
	}

	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		slog.InfoContext(ctx, "gracefully shutting down")
		if err := a.Stop(shutdownCtx); err != nil {
			slog.ErrorContext(ctx, "failed to gracefully shutdown", "error", err)
		}
	}()

	if err := a.Serve(ctx); err != nil {
		slog.ErrorContext(ctx, "cannot serve", "error", err)
	}
}
