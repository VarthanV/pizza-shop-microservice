package cooks

import (
	"context"

	"github.com/VarthanV/kitchen/cooks/models"
)

type Repository interface {
	GetCookByID(ctx context.Context, id int) *models.Cook
	GetAvailableCooks(ctx context.Context, IsVegeterian int) *[]models.Cook
	GetFirstAvailableCook(ctx context.Context, IsVegeterian int, cookCh chan *models.Cook)
}
