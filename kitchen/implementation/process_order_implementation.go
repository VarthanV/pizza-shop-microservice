package implementation

import (
	"context"
	"time"

	"github.com/VarthanV/kitchen/cooks"
	"github.com/VarthanV/kitchen/processes"
	"github.com/VarthanV/kitchen/queue"
	"github.com/golang/glog"
)

type processorderimplementation struct {
	cookservice cooks.Service
	service     processes.OrderProcessUpdateService
}

func NewProcessOrderImplementationService(cs cooks.Service, service processes.OrderProcessUpdateService) processes.OrderProcessService {
	return &processorderimplementation{
		cookservice: cs,
		service:     service,
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
			glog.Errorf("Unable to update the status of cook %s", err)
			// Make
		}
		for _, item := range orderRequest.Details {
			/*
				Just sleeping for 3 seconds to  simulate it as a expensive
				process
			*/
			time.Sleep(3 * time.Second)
			glog.Info("Pizza %s is ready...", item.PizzaID)
			poi.service.MarkOrderItemComplete(ctx, item.PizzaID, orderRequest.OrderUUID)
		}
		poi.service.MarkOrderComplete(ctx, orderRequest.OrderUUID, cookID)
		glog.Info("Trying to update status of the order")
		// Free the cook whether the order fails or not
		defer poi.cookservice.UpdateCookStatus(ctx,cookID,1)
	}()

}
