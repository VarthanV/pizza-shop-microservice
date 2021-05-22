package implementation

import (
	"context"
	"github.com/VarthanV/pizza/pizza/models"
	"github.com/VarthanV/pizza/pizza/services"
)

type orderitemservice struct {
	orderItemRepo models.OrderItemRepository
}

func (o orderitemservice) GetOrderItemByID(ctx context.Context, id int) (*models.OrderItem, error) {
	orderItem, err := o.orderItemRepo.GetOrderItemByID(ctx, id)
	return orderItem, err
}

func (o orderitemservice) GetOrderItemsByOrderID(ctx context.Context, orderID int) (*[]models.OrderItem, error) {
	orderItems, err := o.orderItemRepo.GetOrderItemsByOrderID(ctx, orderID)
	return orderItems, err
}

func (o orderitemservice) AddOrderItem(ctx context.Context, pizzaID int, orderUUID string, Quantity int, Price int) error {
	err := o.orderItemRepo.AddOrderItem(ctx, pizzaID, orderUUID, Quantity, Price)
	return err
}

func NewOrderItemService(repo models.OrderItemRepository) services.OrderItemService {
	return &orderitemservice{
		repo,
	}
}
