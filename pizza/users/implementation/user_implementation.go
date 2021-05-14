package implementation

import (
	"context"
	"errors"

	"github.com/VarthanV/pizza/users"
	"github.com/VarthanV/pizza/users/models"
)

var ErrUnableToGetUser = errors.New("unable to Get user")

type service struct {
	dbRepository models.UserRepository
}

func NewService(repo models.UserRepository) users.Service {
	return &service{
		dbRepository: repo,
	}
}

func (s service) CreateUser(ctx context.Context, user models.User) error {

	return nil
}

func (s service) GetUserById(ctx context.Context, id string) (user models.User, err error) {

	return models.User{}, ErrUnableToGetUser
}

func (s service) LoginUser(ctx context.Context, email string ,password string)  {

}