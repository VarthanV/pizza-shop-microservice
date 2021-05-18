package models

import "context"

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"_"`
	PhoneNumber string `json:"phone_number"`
}

type UserRepository interface {
	CreateUser(ctx context.Context, user User) error
	GetUserById(ctx context.Context, id string) (User, error)
	LoginUser(ctx context.Context, email string, password string) (token TokenDetails)
	GetUserByEmail(ctx context.Context, email string) *User
	GetUserByPhoneNumberOrEmail(ctx context.Context, email string, phoneNumber string) *User
}
