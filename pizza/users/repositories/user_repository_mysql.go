package repositories


import (
	"context"
	"database/sql"
	"errors"
	"github.com/VarthanV/pizza/users/models"
)

var ErrUnableToGetUser = errors.New("unable to Get user")

type mysqlrepository struct {
	db *sql.DB
}

func NewMySqlRepository(db *sql.DB) models.UserRepository {
	return &mysqlrepository{
		db: db,
	}
}

func (r mysqlrepository) CreateUser(ctx context.Context, user models.User) error {
	
	return nil
}

func (r mysqlrepository) GetUserById(ctx context.Context, id string) (models.User, error) {

	return models.User{}, ErrUnableToGetUser
}

func (r mysqlrepository) LoginUser(ctx context.Context, email string, password string) (Token string) {

	return ""
}
