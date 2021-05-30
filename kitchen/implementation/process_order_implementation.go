package implementation

import (
	"context"

	"github.com/VarthanV/kitchen/cooks"
	"github.com/VarthanV/kitchen/processes"
	"github.com/VarthanV/kitchen/queue"
	"github.com/golang/glog"
)

type processorderimplementation struct {
	processOrderRepo processes.OrderProcessRepository
	cookservice      cooks.Service
}

func NewProcessOrderImplementationService(pro processes.OrderProcessRepository, cs cooks.Service) processes.OrderProcessService {
	return &processorderimplementation{
		processOrderRepo: pro,
		cookservice:      cs,
	}
}

func (poi processorderimplementation) ProcessOrder(ctx context.Context, orderRequest queue.OrderQueueRequest, cookID int) {
	/*
		1) Mark the cook as not available.
		2) Loop through each order details.
		3) Process each pizza.
		4) Mark the details in DB.
		5) If all the items are processed sucessfully  mark the order status in DB.

	*/

	var err error
	go func() {
		err = poi.cookservice.UpdateCookStatus(ctx, cookID, 0)
		if err != nil {
			glog.Errorf("Unable to update the status of order %s", err)
			// Make
		}
	}()
	glog.Info("Trying to update status of the order")

}
