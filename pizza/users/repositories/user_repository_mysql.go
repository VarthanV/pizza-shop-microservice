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

func NewMySqlRepository(db *sql.DB) models.UserRepository {
	return &mysqlrepository{
		db: db,
	}
}

func (r mysqlrepository) CreateUser(ctx context.Context, user models.User) error {
	sql := `
		INSERT INTO users (id ,name ,email,password,phone_number)
		VALUES ($1,$2,$3,$4,$5)
	`
	_, err := r.db.Exec(sql, user.ID, user.Name, user.Email, user.Password, user.PhoneNumber)
	if err != nil {
		glog.Error("Unable to insert user with the email", user.Email)
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

func (r mysqlrepository) GetUserByEmail(ctx context.Context, email string) (user models.User) {
	var rowUser models.User
	sql := `
		SELECT * from users
		WHERE email = $1
	`
	rows, err := r.db.Query(sql, email)
	if err != nil {
		glog.Info("Error while fetching data from database")
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&rowUser)
		if err != nil {
			glog.Error("Unable to scan rowss....", err)
		}
	}
	return rowUser
}
