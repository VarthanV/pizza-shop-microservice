package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/VarthanV/pizza/users/models"
	"github.com/golang/glog"
)

var ErrUnableToGetUser = errors.New("unable to Get user")

type mysqlrepository struct {
	db *sql.DB
}

func NewMySqlRepository(db *sql.DB) (repo models.UserRepository, err error) {
	return &mysqlrepository{
		db: db,
	}, nil
}

func (r mysqlrepository) CreateUser(ctx context.Context, user models.User) error {
	sql := `
		INSERT INTO users (id ,name ,email,password,phone_number)
		VALUES (?,?,?,?,?)
	`
	_, err := r.db.ExecContext(ctx, sql, user.ID, user.Name, user.Email, user.Password, user.PhoneNumber)
	if err != nil {
		glog.Error("Unable to insert user.... ", err)
		return err
	}
	return nil
}

func (r mysqlrepository) GetUserById(ctx context.Context, id string) (models.User, error) {

	return models.User{}, ErrUnableToGetUser
}

func (r mysqlrepository) LoginUser(ctx context.Context, email string, password string) (token models.TokenDetails) {

	return models.TokenDetails{}
}

func (r mysqlrepository) GetUserByEmail(ctx context.Context, email string) *models.User {
	var rowUser models.User
	sql := `
		SELECT * from users u
		WHERE email = ?
	`
	err := r.db.QueryRowContext(ctx, sql, email).Scan(&rowUser.ID, &rowUser.Name, &rowUser.Email, &rowUser.Password, &rowUser.PhoneNumber)
	if err != nil {
		glog.Error("Unable to query from user table %s", err)
		return nil
	}
	return &rowUser
}

func (r mysqlrepository) GetUserByPhoneNumberOrEmail(ctx context.Context, email string, phoneNumber string) *models.User {
	var user models.User
	sql := `
		SELECT email, phone_number from users u
		WHERE email = ?
		AND phone_number = ?
	`
	err := r.db.QueryRowContext(ctx, sql, email, phoneNumber).Scan(&user.Email, &user.PhoneNumber)
	if err != nil {
		glog.Errorf("Unable to query from user table %s ....", err)
		return nil
	}
	return &user

}
