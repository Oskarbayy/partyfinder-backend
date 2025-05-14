package inmemory

import (
	"context"
	"fmt"
	"sync"

	"github.com/Oskarbayy/partyfinder-backend/pkg/cqrs"
)

type QueryBus struct {
	handlers   map[string]cqrs.QueryHandler[any]
	middleware []QueryMiddleware
	mu         sync.RWMutex
}

type QueryMiddleware func(next cqrs.QueryHandler[any]) cqrs.QueryHandler[any]

func NewQueryBus(middleware ...QueryMiddleware) *QueryBus {
	return &QueryBus{
		handlers:   make(map[string]cqrs.QueryHandler[any]),
		middleware: middleware, // middleware is implicit conversion to an array
	}
}

func (b *QueryBus) Register(QueryType string, QueryHandler cqrs.QueryHandler[any]) {
	b.mu.Lock()
	defer b.mu.Unlock()

	h := QueryHandler

	for i := len(b.middleware) - 1; i >= 0; i-- {
		h = b.middleware[i](h)
	}

	b.handlers[QueryType] = h
}

func (b *QueryBus) Dispatch(ctx context.Context, Query cqrs.Query) (any, error) {
	b.mu.RLock()
	queryHandler, exists := b.handlers[Query.QueryType()]
	b.mu.RUnlock()

	if !exists {
		return nil, fmt.Errorf("No Query handler found for QueryType: %q", Query.QueryType())
	}

	return queryHandler.Handle(ctx, Query)
}
