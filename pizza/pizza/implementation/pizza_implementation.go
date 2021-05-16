package implementation

import (
	"context"

	"github.com/VarthanV/pizza/pizza"
	"github.com/VarthanV/pizza/pizza/models"
)

type pizzaservice struct {
	dbRepository models.PizzaRepository
}

func NewService(repo models.PizzaRepository) pizza.Service {
	return  &pizzaservice{
		dbRepository: repo,
	}
}

func (s pizzaservice) GetAllPizzas(ctx context.Context, isVeg int) (pizza[]models.Pizza, err error){
	pizzas ,err:= s.dbRepository.GetAllPizzas(ctx, isVeg)
	return pizzas,err
}