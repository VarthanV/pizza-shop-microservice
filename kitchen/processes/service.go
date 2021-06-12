package processes

import (
	"context"

	"github.com/VarthanV/kitchen/queue"
)

type OrderProcessUpdateService interface {
	MarkOrderComplete(ctx context.Context, orderUUID string, cookID int) error
	// Updating status of  a single pizza item
	MarkOrderItemComplete(ctx context.Context, pizzaID int, orderUUID string) error
}

type OrderProcessUpdateRepo interface {
	UpdateOrderProcces(ctx context.Context, orderUUID string, cookID int) error
	// Updating status of  a single pizza item
	UpdateOrderItemProcess(ctx context.Context, pizzaID int, orderUUID string) error
}

type OrderRequestService interface {
	SubmitOrderRequest(ctx context.Context, request queue.OrderQueueRequest, c chan bool)
}

type OrderProcessService interface {
	ProcessOrder(ctx context.Context, orderRequest queue.OrderQueueRequest, cookID int, updateStatus bool)
}
