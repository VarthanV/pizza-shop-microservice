package pizza

import (
	"context"

	"github.com/VarthanV/pizza/pizza/models"
)

type Service interface {
	GetAllPizzas(ctx context.Context, isVeg int) (pizza []models.Pizza ,err error)
}
