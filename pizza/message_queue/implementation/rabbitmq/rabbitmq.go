package rabbitmq

import (
	"context"
	"encoding/json"

	"github.com/VarthanV/pizza/message_queue"
	"github.com/VarthanV/pizza/pizza"
	"github.com/golang/glog"
)

type rabbitmqImplementation struct {
	repo               message_queue.QueueRepository
	orderupdateservice pizza.OrderUpdateService
}

func NewRabbitMQService(repo message_queue.QueueRepository, ousvc pizza.OrderUpdateService) message_queue.QueueService {
	return rabbitmqImplementation{
		repo:               repo,
		orderupdateservice: ousvc,
	}
}

func (r rabbitmqImplementation) PublishOrderDetails(ctx context.Context, request message_queue.OrderQueueRequest) error {
	glog.Infof("Posting message to the Queue...")
	err := r.repo.PublishOrderDetails(ctx, request)
	return err
}

func (r rabbitmqImplementation) ConsumeOrderStatus(ctx context.Context) {
	glog.Infof("Consuming message from the Message queue...")
	msgs, err := r.repo.ConsumeOrderStatus(ctx)
	if err != nil {
		glog.Errorf("Unable to consume messages from the queue")
	}
	go func() {
		for msg := range msgs {
			var orderUpdateMsg message_queue.OrderUpdateStatusRequest
			glog.Infof(string(msg.Body))
			err := json.Unmarshal(msg.Body, &orderUpdateMsg)
			if err != nil {
				glog.Errorf("Unable to unmarshal request from message queue %s ..", err)
				return
			}
			err = r.orderupdateservice.UpdateOrderStatus(orderUpdateMsg.OrderUUID, orderUpdateMsg.Status)
			if err != nil {
				glog.Errorf("Unable to update the status of the order.. %s", err)
			}
		}
	}()
}
