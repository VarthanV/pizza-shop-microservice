package implementation

import (
	"context"

	"github.com/VarthanV/pizza/pizza"
	"github.com/VarthanV/pizza/pizza/models"
)

type service struct {
	dbRepository models.PizzaRepository
}

func NewService(repo models.PizzaRepository) pizza.Service {
	return  &service{
		dbRepository: repo,
	}
}

func (s service) GetAllPizzas(ctx context.Context, isVeg int) (pizza[]models.Pizza, err error){
	pizzas ,err:= s.dbRepository.GetAllPizzas(ctx, isVeg)
	return pizzas,err
}