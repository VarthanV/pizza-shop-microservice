package implementation

import (
	"context"
	"errors"

	"github.com/VarthanV/pizza/pizza"
)

type orderupdateimplementation struct {
	repo pizza.OrderUpdateRepo
}

func NewOrderUpdateImplementation(repo pizza.OrderUpdateRepo) pizza.OrderUpdateService {
	return &orderupdateimplementation{
		repo: repo,
	}
}

func (ou orderupdateimplementation) UpdateOrderStatus(orderUUID string, status string) error {
	if orderUUID == "" {
		return errors.New("Order id  cannot be empty")
	}
	ctx := context.Background()
	err := ou.repo.UpdateOrderStatus(ctx, orderUUID, status)
	return err
}
