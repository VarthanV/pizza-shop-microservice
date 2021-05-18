package services

import (
	"context"

	"github.com/VarthanV/pizza/pizza/models"
)

type CartService interface {
	GetCart(ctx context.Context, userId string) (*[]models.CartQueryResult, error)
	AddItem(ctx context.Context, itemId int, userId string, quantity int, price int) error
	EditItem(ctx context.Context, cartItemId int, itemId int, quantity int, price int) error
	DeleteItem(ctx context.Context, cartItemId int, userId string) error
}
