package services

import (
	"context"
	"github.com/VarthanV/pizza/pizza/models"
)

type OrderItemService interface {
	GetOrderItemByID(ctx context.Context, id int) (*models.OrderItem, error)
	GetOrderItemsByOrderID(ctx context.Context, orderID int) (*[]models.OrderItem, error)
	AddOrderItem(ctx context.Context, pizzaID int, orderUUID string, Quantity int, Price int) error
}
