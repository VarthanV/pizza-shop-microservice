package implementation

import (
	"context"
	"time"

	"github.com/VarthanV/kitchen/processes"
	"github.com/VarthanV/kitchen/queue"
	"github.com/golang/glog"
)

type orderrequestprocessingimplementation struct {
	service processes.OrderProcessUpdateService
}

func NewProcessOrderRepository(svc processes.OrderProcessUpdateService) processes.OrderProcessService {
	return &orderrequestprocessingimplementation{
		service: svc,
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
			op.service.MarkOrderItemComplete(ctx, item.PizzaID, request.OrderUUID)
		}
		op.service.MarkOrderComplete(ctx, request.OrderUUID, cookID)
	}()
}
