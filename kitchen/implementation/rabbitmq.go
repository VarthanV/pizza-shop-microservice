package implementation

import (
	"context"
	"encoding/json"
	"time"

	"github.com/VarthanV/kitchen/processes"
	"github.com/VarthanV/kitchen/queue"
	"github.com/VarthanV/kitchen/shared"
	"github.com/golang/glog"
)

type rmqimplementation struct {
	repo                queue.QueueRepository
	orderRequestService processes.OrderRequestService
}

func NewRabbitMQService(repo queue.QueueRepository, orc processes.OrderRequestService) queue.QueueService {
	return rmqimplementation{
		repo:                repo,
		orderRequestService: orc,
	}
}

func (r rmqimplementation) PublishOrderStatus(ctx context.Context) error {
	glog.Infof("Posting message to the Queue...")
	err := r.repo.PublishOrderStatus(ctx)
	return err

}

func (r rmqimplementation) ConsumeOrderDetails(ctx context.Context) {
	glog.Infof("Consuming message from the Message queue...")
	msgs, err := r.repo.ConsumeOrderDetails(ctx)
	if err != nil {
		glog.Errorf("Unable to consume messages from the queue")
	}
	go func() {
		for msg := range msgs {
			var req queue.OrderQueueRequest
			m := msg
			err := json.Unmarshal(m.Body, &req)
			if err != nil {
				glog.Error("Unable to unmarshal json.. ", err)
			}
			glog.Info("Consumed message .. ", req)
			ctx, cancel := context.WithDeadline(ctx, time.Now().Add(shared.DeadlineForOrderSubmitRequest))
			c := make(chan bool)
			go r.orderRequestService.SubmitOrderRequest(ctx, req, c)
			select {
			case consumed := <-c:
				/*
					If the kitchen has cooks who are free
					we will send a response via this channel and update
					the order status here itself if the value is true
					else it will be updated when the cook takes up
				*/
				glog.Info("Consumed..", consumed)
			case <-ctx.Done():
				glog.Info("The context is done..")
			}
			defer cancel()
		}
	}()
}
