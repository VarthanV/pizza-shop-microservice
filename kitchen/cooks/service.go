package cooks

import (
	"context"

	"github.com/VarthanV/kitchen/cooks/models"
)

type Service interface {
	GetCookByID(ctx context.Context, id int) *models.Cook
	GetAvailableCooks(ctx context.Context) *[]models.Cook
	GetFirstAvailableCook(ctx context.Context, ch chan *models.Cook)
	UpdateCookStatus(ctx context.Context, cookID int, status int) error
}
