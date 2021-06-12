package implementation

import (
	"context"
	"encoding/json"

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

func (r rmqimplementation) PublishOrderStatus(ctx context.Context, orderUUID string, status string) error {
	glog.Infof("Posting message to the Queue...")
	err := r.repo.PublishOrderStatus(ctx, orderUUID, status)
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
				ctx.Done()
				return
			}
			glog.Info("Consumed message .. ", req.OrderUUID)
			ctx := context.Background()
			c := make(chan bool)
			go r.orderRequestService.SubmitOrderRequest(ctx, req, c)
			select {
			case isCookAvailable := <-c:
				/*
					If the kitchen has cooks who are free
					we will send a response via this channel and update
					the order status here itself if the value is true
					else it will be updated when the cook takes up
				*/
				glog.Info("Consumed..", isCookAvailable)
				if isCookAvailable == true {
					r.PublishOrderStatus(ctx, req.OrderUUID, shared.OrderStatusProcessing)
				} else {
					glog.Info("Cook is not available will be updated once the order is process and complete")
					r.PublishOrderStatus(ctx,req.OrderUUID,shared.OrderStatusWaitingForCook)
				}
			}
		}
	}()
}
