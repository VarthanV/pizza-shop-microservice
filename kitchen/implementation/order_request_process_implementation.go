package implementation

import (
	"context"
	"time"

	"github.com/VarthanV/kitchen/processes"
	"github.com/VarthanV/kitchen/queue"
	"github.com/golang/glog"
)

type orderrequestprocessingimplementation struct {
	processOrderRepo processes.OrderProcessRepository
	
}

func NewProcessOrderRepository(repo processes.OrderProcessRepository) processes.OrderProcessService {
	return &orderrequestprocessingimplementation{
		processOrderRepo: repo,
	}
}

func (op orderrequestprocessingimplementation) ProcessOrder(ctx context.Context, request queue.OrderQueueRequest, cookID int) {
	go func() {
		for _, item := range request.Details {
			/*
				Just sleeping for 3 seconds to  simulate it as a expensive
				process
			*/
			time.Sleep(3 * time.Second)
			glog.Info("Pizza %s is ready...", item.PizzaID)
		}
	}()
}
