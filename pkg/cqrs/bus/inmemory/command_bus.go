package inmemory

import (
	"context"
	"fmt"
	"sync"

	"github.com/Oskarbayy/partyfinder-backend/pkg/cqrs"
)

type CommandBus struct {
	handlers   map[string]cqrs.CommandHandler
	middleware []CommandMiddleware
	mu         sync.RWMutex
}

type CommandMiddleware func(next cqrs.CommandHandler) cqrs.CommandHandler

func NewCommandBus(middleware ...CommandMiddleware) *CommandBus {
	return &CommandBus{
		handlers:   make(map[string]cqrs.CommandHandler),
		middleware: middleware, // middleware is implicit conversion to an array
	}
}

func (b *CommandBus) Register(commandType string, commandHandler cqrs.CommandHandler) {
	b.mu.Lock()
	defer b.mu.Unlock()

	h := commandHandler

	for i := len(b.middleware) - 1; i >= 0; i-- {
		h = b.middleware[i](h)
	}

	b.handlers[commandType] = h
}

func (b *CommandBus) Dispatch(ctx context.Context, command cqrs.Command) error {
	b.mu.RLock()
	cmdHandler, exists := b.handlers[command.CommandType()]
	b.mu.RUnlock()

	if !exists {
		return fmt.Errorf("No command handler found for commandType: %q", command.CommandType())
	}

	return cmdHandler.Handle(ctx, command)
}
