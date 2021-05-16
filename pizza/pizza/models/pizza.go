package models

import "context"

type Pizza struct {
	ID           int
	Name         string
	Price        uint
	IsVegeterian int
}

type PizzaRepository interface {
	GetAllPizzas(ctx context.Context, isVegetarian int) (pizzas []Pizza,err error)
	// GetPizzaInPriceRangeLessThan(ctx context.Context, price uint) (pizzas []Pizza)
	// GetPizzaInPriceRangeGreaterThan(ctx context.Context, price int) (pizzas []Pizza)
}
