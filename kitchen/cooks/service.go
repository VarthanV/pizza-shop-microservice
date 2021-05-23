package cooks

import (
	"context"

	"github.com/VarthanV/kitchen/cooks/models"
)

type Service interface{
	GetCookByID(ctx context.Context,id int) *models.Cook

}