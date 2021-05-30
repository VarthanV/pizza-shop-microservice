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

func (c cookservice) GetAvailableCooks(ctx context.Context, IsVegeterian int) *[]models.Cook {
	cooks := c.cookRepo.GetAvailableCooks(ctx, IsVegeterian)
	return cooks
}

func (c cookservice) GetFirstAvailableCook(ctx context.Context, IsVegeterian int , ch chan *models.Cook) {
	cookCh := make(chan *models.Cook, 1)
	go func() {
		c.cookRepo.GetFirstAvailableCook(ctx, IsVegeterian,cookCh)
		select {
		case cook := <-cookCh:
			ch <- cook
			close(ch)
		}
	}()
}
