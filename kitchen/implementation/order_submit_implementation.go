package implementation

import (
	"context"
	"time"

	"github.com/VarthanV/kitchen/processes"
	"github.com/VarthanV/kitchen/queue"
)

type ordersubmitimplementation struct {
}

func NewOrderRequestImplementation() processes.OrderRequestService {
	return &ordersubmitimplementation{}
}

func (op ordersubmitimplementation) SubmitOrderRequest(ctx context.Context, request queue.OrderQueueRequest, c chan bool) {
	time.Sleep(time.Second * 20)
	c <- false
	return
}
