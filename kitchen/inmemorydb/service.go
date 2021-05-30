package inmemorydb

import (
	"context"

	"github.com/VarthanV/kitchen/queue"
)

type OrderRequestInMemoryService interface {
	SetOrder(ctx context.Context, key string, request queue.OrderQueueRequest) error
	GetOrder(ctx context.Context, key string) (*queue.OrderQueueRequest, error)
}
