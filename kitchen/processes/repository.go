package processes

import (
	"context"
)

type OrderProcessRepository interface {
	CompleteOrder(ctx context.Context, orderID int, cookID int) error
}
