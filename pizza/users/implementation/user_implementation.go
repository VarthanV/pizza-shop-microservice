package implementation

import (
	"context"
	"errors"

	"github.com/VarthanV/pizza/users"
	"github.com/VarthanV/pizza/users/models"
	"github.com/VarthanV/pizza/users/utils"
	"github.com/golang/glog"
	"github.com/google/uuid"
)

var ErrUnableToGetUser = errors.New("unable to Get user")

type service struct {
	dbRepository    models.UserRepository
	tokenRepository models.TokenRepository
}

func NewService(repo models.UserRepository, tokenRepo models.TokenRepository) users.Service {
	return &service{
		dbRepository:    repo,
		tokenRepository: tokenRepo,
	}
}

func (s service) CreateUser(ctx context.Context, user models.User) error {
	rowUser := s.dbRepository.GetUserByEmail(ctx, user.Email)

	if rowUser != nil {
		glog.Info("User exists with the same email...")
		return  errors.New("conflict")
	}
	// Do some cleanup
	user.ID = uuid.NewString()
	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		glog.Error("Unable to hash password")
		return err
	}
	user.Password = hashed

	//Pass to repository to create user
	err = s.dbRepository.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	glog.Infof("Created user with email %s  successfully",user.Email)
	return nil
}

func (s service) GetUserById(ctx context.Context, id string) (user models.User, err error) {

	return models.User{}, ErrUnableToGetUser
}

func (s service) LoginUser(ctx context.Context, email string, password string) {

}
