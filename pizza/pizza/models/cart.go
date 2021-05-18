package models

import "context"

type Cart struct {
	ID       int
	UserID   string
	PizzaID  int
	Quantity int
	Price    int
}

type CartQueryResult struct {
	PizzaName    string `json:"pizza_name"`
	Price        int `json:"price"`
	Quantity     int `json:"quantity"`
	IsVegeterian int `json:"is_vegeterian"`
}

type CartRepository interface {
	GetCart(ctx context.Context, userId string) (*[]CartQueryResult, error)
	AddItem(ctx context.Context, itemId int, userId string, quantity int, price int) error
	EditItem(ctx context.Context, cartItemId int, itemId int, quantity int, price int) error
	DeleteItem(ctx context.Context, cartItemId int, userId string) error
	GetCartItem(ctx context.Context, itemId int, userId string) *Cart
}
