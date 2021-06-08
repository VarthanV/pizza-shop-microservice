package pizza

import "context"

type OrderUpdateService interface {
	UpdateOrderStatus(orderUUID string, status string) error
}

type OrderUpdateRepo interface {
	UpdateOrderStatus(ctx context.Context, orderUUID string, status string) error
}