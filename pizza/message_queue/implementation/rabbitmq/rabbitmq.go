package rabbitmq

import (
	"context"
	"github.com/VarthanV/pizza/message_queue"
	"github.com/golang/glog"
)

type rabbitmqImplementation struct {
	repo message_queue.QueueRepository
}



func NewRabbitMQService(repo message_queue.QueueRepository)  message_queue.QueueService{
	return rabbitmqImplementation{
		repo: repo,
	}
}

func (r rabbitmqImplementation) PublishOrderDetails(ctx context.Context, request message_queue.OrderQueueRequest) error {
	glog.Infof("Posting message to the Queue...")
	err := r.repo.PublishOrderDetails(ctx,request)
	return err
}

func (r rabbitmqImplementation) ConsumeOrderStatus(ctx context.Context) {
	glog.Infof("Consuming message from the Message queue...")
	msgs ,err  := r.repo.ConsumeOrderStatus(ctx)
	if err != nil {
		glog.Errorf("Unable to consume messages from the queue")
	}
	go func() {
		for msg := range  msgs{
			glog.Infof(string(msg.Body))
		}
	}()
}