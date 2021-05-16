package models

import "context"

type Pizza struct {
	ID           int    `json:"int"`
	Name         string `json:"name"`
	Price        uint   `json:"price"`
	IsVegeterian int    `json:"is_vegeterian"`
}

type PizzaRepository interface {
	GetAllPizzas(ctx context.Context, isVegetarian int) (pizzas []Pizza, err error)
	// GetPizzaInPriceRangeLessThan(ctx context.Context, price uint) (pizzas []Pizza)
	// GetPizzaInPriceRangeGreaterThan(ctx context.Context, price int) (pizzas []Pizza)
}
