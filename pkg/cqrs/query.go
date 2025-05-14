package cqrs

import "context"

type Query interface {
	QueryType() string
}

type QueryHandler[T any] interface {
	Handle(ctx context.Context, query Query) (T, error)
}

type QueryBus interface {
	Dispatch(ctx context.Context, query Query) (any, error)
	Register(queryType string, handler QueryHandler[any])
}
