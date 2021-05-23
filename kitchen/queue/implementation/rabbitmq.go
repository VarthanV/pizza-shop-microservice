package implementation

import (
	"context"
	"encoding/json"


	"github.com/VarthanV/kitchen/queue"
	"github.com/golang/glog"
)

type rmqimplementation struct {
	repo queue.QueueRepository 
}



func NewRabbitMQService(repo queue.QueueRepository)  queue.QueueService{
	return rmqimplementation{
		repo: repo,
	}
}



func (r rmqimplementation) PublishOrderStatus(ctx context.Context) error {
	glog.Infof("Posting message to the Queue...")
	err := r.repo.PublishOrderStatus(ctx)
	return err

}

func (r rmqimplementation)  ConsumeOrderDetails(ctx context.Context) { 
	glog.Infof("Consuming message from the Message queue...")
	msgs ,err  := r.repo.ConsumeOrderDetails(ctx)
	if err != nil {
		glog.Errorf("Unable to consume messages from the queue")
	}
	go func() {
		for msg := range  msgs{
			var req queue.OrderQueueRequest
			err := json.Unmarshal(msg.Body,&req)
			if err !=nil {
				glog.Error("Unable to unmarshal json.. ",err)
			}
			glog.Info("Consumed message .. ",req)
		}
	}()
	
}