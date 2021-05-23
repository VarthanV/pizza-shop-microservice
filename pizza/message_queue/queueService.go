package message_queue

import (
	"context"
	"github.com/streadway/amqp"
)

type QueueRepository interface {
	PublishOrderDetails(ctx context.Context,request OrderQueueRequest) error
	ConsumeOrderStatus(ctx context.Context) (<- chan amqp.Delivery,error)
}

type QueueService interface {
	PublishOrderDetails(ctx context.Context,request OrderQueueRequest) error
	ConsumeOrderStatus(ctx context.Context)
}
