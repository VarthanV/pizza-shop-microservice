package implementation

import (
	"context"
	"errors"
	"github.com/VarthanV/pizza/pizza/models"
	"github.com/VarthanV/pizza/pizza/services"
	"github.com/golang/glog"
)

type cartservice struct {
	cartRepo models.CartRepository
}

func NewCartService(repo models.CartRepository) services.CartService {

	return &cartservice{
		cartRepo: repo,
	}
}

func (c cartservice) GetCart(ctx context.Context, userId string) (*[]models.CartQueryResult, error) {
	if userId == "" {
		glog.Error("Cannot get cart for empty user")
		return nil, errors.New("cannot get cart for empty user")
	}
	cart, err := c.cartRepo.GetCart(ctx, userId)
	if err != nil {
		glog.Errorf("Error getting cart %s", err)
		return nil, err
	}
	return cart, err
}

func (c cartservice) AddItem(ctx context.Context, itemId int, userId string, quantity int, price int) error {
	//See first if an item  is there in the users cart already if so return a response conflict
	item := c.cartRepo.GetCartItem(ctx, itemId, userId)
	if item != nil {
		glog.Errorf("The item already exists in the users cart")
		return errors.New("item-conflict")
	}
	err := c.cartRepo.AddItem(ctx, itemId, userId, quantity, price)
	if err != nil {
		glog.Errorf("Error adding item to cart %s", err)
		return err
	}
	return nil
}

func (c cartservice) EditItem(ctx context.Context, cartItemId int, itemId int, quantity int, price int) error {
	err := c.cartRepo.EditItem(ctx, cartItemId, itemId, quantity, price)
	if err != nil {
		glog.Errorf("Error updating cart %s", err)
		return err
	}
	return nil
}

func (c cartservice) DeleteItem(ctx context.Context, cartItemId int, userId string) error {
	err := c.cartRepo.DeleteItem(ctx, cartItemId, userId)
	if err != nil {
		return err
	}
	return nil
}
