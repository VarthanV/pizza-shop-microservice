package handlers

import "errors"

func (c *AddToCartRequest) Validate() error {
	if c.Quantity == 0 {
		return errors.New("enter a valid quantity")
	}
	if c.PizzaID == 0 {
		return errors.New("enter a valid item id")
	}
	if c.Price == 0 {
		return errors.New("enter a valid price")
	}
	return nil
}

func (c *EditCartRequest) Validate() error {
	if c.ID == 0 {
		return errors.New("enter a valid id")
	}
	if err := c.AddToCartRequest.Validate(); err != nil {
		return err
	}
	return nil
}
