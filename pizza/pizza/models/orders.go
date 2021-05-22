package models

import "context"

type Order struct {
	ID          int
	OrderUUID   string
	UserID      string
	Total       int
	OrderStatus string
}

type OrderRepository interface {
	CreateOrder(ctx context.Context, order Order, userID string, cart *[]CartQueryResult) (err error)
	GetOrderByUUID(ctx context.Context, uuid string) (*Order, error)
	GetOrdersByUserID(ctx context.Context, userId int) (*[]Order, error)
}
