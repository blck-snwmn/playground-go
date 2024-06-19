package ctxutil

import "context"

type key[T any] struct{}

func WithValue[T any](ctx context.Context, val T) context.Context {
	return context.WithValue(ctx, key[T]{}, val)
}

func Value[T any](ctx context.Context) T {
	return ctx.Value(key[T]{}).(T)
}
