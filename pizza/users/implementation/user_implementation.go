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
	dbRepository models.UserRepository
	tokenService users.TokenService
	utils        utils.UtilityService
}

func NewService(repo models.UserRepository, tokenService users.TokenService ,utilityservice utils.UtilityService) users.Service {
	return &service{
		dbRepository: repo,
		tokenService: tokenService,
		utils: utilityservice,
	}
}

func (s service) CreateUser(ctx context.Context, user models.User) error {
	rowUser := s.dbRepository.GetUserByPhoneNumberOrEmail(ctx, user.Email, user.PhoneNumber)

	if rowUser != nil {
		glog.Info("User exists with the same email or phone number...")
		return errors.New("conflict")
	}
	// Do some cleanup
	user.ID = uuid.NewString()
	hashed, err := s.utils.HashPassword(user.Password)
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
	glog.Infof("Created user with email %s  successfully", user.Email)
	return nil
}

func (s service) GetUserById(ctx context.Context, id string) (user models.User, err error) {

	return models.User{}, ErrUnableToGetUser
}

func (s service) LoginUser(ctx context.Context, email string, password string) (*models.TokenDetails, error) {
	user := s.dbRepository.GetUserByEmail(ctx, email)
	// Compare password
	isValid := s.utils.CheckPasswordHash(password, user.Password)
	if !isValid {
		return nil, errors.New("login credentials invalid")
	}
	tokenDetails, err := s.tokenService.CreateToken(ctx, *user)
	if err != nil {
		glog.Errorf("Unable to generate token for the user... %f", err)
	}
	return &tokenDetails, nil
}
