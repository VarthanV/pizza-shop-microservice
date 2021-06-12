package redisclient

import (
	"context"

	"github.com/VarthanV/kitchen/inmemorydb"
	"github.com/go-redis/redis/v8"
	"github.com/golang/glog"
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
	err := oq.client.Set(ctx, "orders", request, 0).Err()
	glog.Error(err)
	return err
}

func (oq orderqueueredis) GetOrder(ctx context.Context, key string) string {
	return ""
}
