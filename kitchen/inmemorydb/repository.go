package inmemorydb

import "context"

type OrderRequestInMemoryRepo interface {
	SetOrder(ctx context.Context, key string, request string) error
	GetOrder(ctx context.Context, key string) (string ,error)
}
	