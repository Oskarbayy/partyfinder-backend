package cqrs

import "context"

type Command interface {
	CommandType() string
}

type CommandHandler interface {
	Handle(ctx context.Context, cmd Command) error
}

type CommandBus interface {
	Dispatch(ctx context.Context, cmd Command) error
	Register(cmdType string, handler CommandHandler)
}
