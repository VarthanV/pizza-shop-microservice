package implementation

import (
	"context"

	"github.com/VarthanV/kitchen/cooks"
	"github.com/VarthanV/kitchen/cooks/models"
	"github.com/golang/glog"
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

func (c cookservice) GetAvailableCooks(ctx context.Context) *[]models.Cook {
	cooks := c.cookRepo.GetAvailableCooks(ctx)
	return cooks
}

func (c cookservice) GetFirstAvailableCook(ctx context.Context, ch chan *models.Cook) {
	cookCh := make(chan *models.Cook, 1)
	go func() {
		c.cookRepo.GetFirstAvailableCook(ctx, cookCh)
		select {
		case cook := <-cookCh:
			ch <- cook
			close(ch)
		}
	}()
}

func (c cookservice) UpdateCookStatus(ctx context.Context, cookID int, status int) error {
	err := c.cookRepo.UpdateCookStatus(ctx, cookID, status)
	if err != nil {
		glog.Info("Updated the cook status , So marking it as done")
		ctx.Done()
	}
	return err
}
