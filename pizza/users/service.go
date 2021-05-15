package users

import (
	"context"

	"github.com/VarthanV/pizza/users/models"
)

// This interface defines the method this package will expose
type Service interface{
	CreateUser(ctx context.Context, user models.User) error
	GetUserById(ctx context.Context, id string) (user models.User, err error)
	LoginUser(ctx context.Context, email string, password string) (*models.TokenDetails, error)
}