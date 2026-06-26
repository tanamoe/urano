package hooks

import (
	"context"
	"errors"
	"sync"
)

type Handler[T any] struct {
	Func func(context.Context, T) error
}

type Hook[T any] struct {
	handlers []*Handler[T]
	mu       sync.RWMutex
}

func (h *Hook[T]) BindFunc(fn func(context.Context, T) error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.handlers = append(h.handlers, &Handler[T]{Func: fn})
}

func (h *Hook[T]) Trigger(ctx context.Context, t T) error {
	var errs error
	for _, handler := range h.handlers {
		if err := handler.Func(ctx, t); err != nil {
			errs = errors.Join(errs, err)
		}
	}

	return errs
}
