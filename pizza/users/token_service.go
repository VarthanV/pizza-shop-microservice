package users

import (
	"context"

	"github.com/VarthanV/pizza/users/models"
)

type TokenService interface{
	CreateToken(ctx context.Context, user models.User) (tokenDetails models.TokenDetails, err error) 
	VerifyTokenValidity(ctx context.Context, acToken string) (isValid bool)
}