package cooks

import (
	"context"

	"github.com/VarthanV/kitchen/cooks/models"
)

type Repository interface{
	GetCookByID(ctx context.Context,id int) *models.Cook
}