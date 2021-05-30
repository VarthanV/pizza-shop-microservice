package processes

import (
	"context"

	"github.com/VarthanV/kitchen/queue"
)

type PizzaProcessStatusService interface {
}

type OrderProcessUpdateService interface {
	UpdateOrderStatusToLocal(orderUUID string, status string, cookID string) error
	// Updating status of  a single pizza item
	UpdateOrderItemStatus(pizzaID int, status string, cookID string) error
}

type OrderPreparationService interface {
	PrepareOrder(pizzaID string, orderUUID string) error
}

type OrderRequestService interface {
	SubmitOrderRequest(ctx context.Context, request queue.OrderQueueRequest, c chan bool)
}

type OrderProcessService interface {
	ProcessOrder(ctx context.Context, orderRequest queue.OrderQueueRequest, cookID int)
}
