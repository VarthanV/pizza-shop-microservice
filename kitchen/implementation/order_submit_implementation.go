package implementation

import (
	"context"

	"github.com/VarthanV/kitchen/cooks"
	"github.com/VarthanV/kitchen/cooks/models"
	"github.com/VarthanV/kitchen/processes"
	"github.com/VarthanV/kitchen/queue"
	"github.com/golang/glog"
)

type ordersubmitimplementation struct {
	cookservice         cooks.Service
	processOrderService processes.OrderProcessService
}

func NewOrderRequestImplementation(cooksvc cooks.Service, pos processes.OrderProcessService) processes.OrderRequestService {
	return &ordersubmitimplementation{
		cookservice:         cooksvc,
		processOrderService: pos,
	}
}

func (op ordersubmitimplementation) SubmitOrderRequest(ctx context.Context, request queue.OrderQueueRequest, c chan bool) {
	cookChan := make(chan *models.Cook)
	/*
		1) Get list of the available cooks.
		2) If available assign the task to the first available cook.
		3) Send true to the channel so that it can send a true to the channel, so the caller can send
		   order status to the queue.
		4) If no cook present cache the order details in the Redis store and send it first
	*/
	go func() {
		op.cookservice.GetFirstAvailableCook(ctx, cookChan)
		cook := <-cookChan
		glog.Info("Received cook is...", cook)
		if cook != nil {
			c <- true
			close(c)
			/*
				1) Assign the order to the cook
				2) Start a go routine so that the cook can process the order.
				3) Make the cook availability to 0

			*/
			op.processOrderService.ProcessOrder(ctx, request, cook.ID)
			return
		} else {
			c <- false
			close(c)
			return
		}

	}()
	glog.Info("Waiting for the availalbe cook ....")
}
