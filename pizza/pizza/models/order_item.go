package models

import "context"

type OrderItem struct {
	ID       int `json:"id"`
	PizzaID  int `json:"pizza_id"`
	OrderID  int `json:"order_id"`
	Quantity int `json:"quantity"`
	Price    int `json:"price"`
}

//The result which will be sent to frontend will be the join of the pizza table and the item table
//So this is struct is required

type OrderItemQueryResult struct {
	OrderItemPizza Pizza     `json:"pizza"`
	Item           OrderItem `json:"order_item"`
}

type OrderItemRepository interface {
	GetOrderItemByID(ctx context.Context, id int) (*OrderItem, error)
	GetOrderItemsByOrderID(ctx context.Context, orderID int) (*[]OrderItem, error)
	AddOrderItem(ctx context.Context, pizzaID int, orderUUID string, Quantity int, Price int) error
	//	Orders once placed cannot be edited or deleted so no methods exposed for it ,
	//	May be a internal method might exist in future
}
