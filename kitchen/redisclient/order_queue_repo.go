package redisclient

import (
	"context"

	"github.com/VarthanV/kitchen/inmemorydb"
	"github.com/go-redis/redis/v8"
)

type orderqueueredis struct {
	client *redis.Client
}

func NewOrderQueueRepo(client *redis.Client) inmemorydb.OrderRequestInMemoryRepo {
	return &orderqueueredis{
		client: client,
	}
}

func (oq orderqueueredis) SetOrder(ctx context.Context, key string, request string) error {
	err := oq.client.Set(ctx, key, request, 0).Err()
	
	return err
}

func (oq orderqueueredis) GetOrder(ctx context.Context, key string) string {
	return ""
}
