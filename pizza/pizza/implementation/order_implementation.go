package implementation

import (
	"context"

	"github.com/VarthanV/pizza/pizza"
	"github.com/VarthanV/pizza/pizza/models"
	"github.com/google/uuid"
)


type orderservice struct {
	repo models.OrderRepository
}

func NewOrderService(repo models.OrderRepository) pizza.OrderService {
	return &orderservice{
		repo: repo,
	}
}

func (o orderservice) CreateOrder(ctx context.Context, order models.Order) (err error){
	// Assign a uuid to the order
	order.OrderUUID  = uuid.NewString()
	createErr := o.repo.CreateOrder(ctx,order)
	return createErr
}

func (o orderservice) GetOrderByUUID(ctx context.Context,uuid string) (*models.Order, error){
	order ,err := o.repo.GetOrderByUUID(ctx,uuid)
	return order,err
}

func (o orderservice) GetOrdersByUserID(ctx context.Context, userId int) (*[]models.Order, error){
	orders ,err := o.repo.GetOrdersByUserID(ctx,userId)
	return orders,err
}