package processes

import (
	"context"
)

type PizzaProcessRepository interface {
	ProcessPizza(pizzaID int, isVegeterian int) error
}

type OrderProcessRepository interface {
	CompleteOrder(ctx context.Context, orderID int, cookID int) error
}
