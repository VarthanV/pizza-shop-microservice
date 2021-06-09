package queue

import (
	"context"

	"github.com/streadway/amqp"
)

type QueueRepository interface {
	ConsumeOrderDetails(ctx context.Context) (<-chan amqp.Delivery, error)
	PublishOrderStatus(ctx context.Context, orderUUID string, status string) error
}

type QueueService interface {
	ConsumeOrderDetails(ctx context.Context)
	PublishOrderStatus(ctx context.Context, orderUUID string, status string) error
}
