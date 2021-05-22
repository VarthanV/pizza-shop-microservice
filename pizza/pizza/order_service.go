package pizza

import "context"

import "github.com/VarthanV/pizza/pizza/models"

type OrderService interface {
	CreateOrder(ctx context.Context, userID string) (err error)
	GetOrderByUUID(ctx context.Context, uuid string) (*models.Order, error)
	GetOrdersByUserID(ctx context.Context, userId int) (*[]models.Order, error)
}
