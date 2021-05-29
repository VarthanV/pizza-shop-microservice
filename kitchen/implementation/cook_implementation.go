package implementation

import (
	"context"

	"github.com/VarthanV/kitchen/cooks"
	"github.com/VarthanV/kitchen/cooks/models"
)

type cookservice struct {
	cookRepo cooks.Repository
}

func NewCookService(repo cooks.Repository) cooks.Service {
	return cookservice{
		cookRepo: repo,
	}
}

func (c cookservice) GetCookByID(ctx context.Context, id int) *models.Cook {
	cook := c.cookRepo.GetCookByID(ctx, id)
	return cook
}
