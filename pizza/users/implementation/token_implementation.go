package implementation

import (
	"context"

	"github.com/VarthanV/pizza/users"
	"github.com/VarthanV/pizza/users/models"
)

type tokenservice struct {
	tokenRepo models.TokenRepository
}

func NewTokenService(tokenRepo models.TokenRepository) users.TokenService {
	return &tokenservice{
		tokenRepo: tokenRepo,
	}
}

func (t tokenservice) CreateToken(ctx context.Context, user models.User) (tokenDetails models.TokenDetails, err error) {
	token, err := t.tokenRepo.CreateToken(ctx, user)
	return token, err
}

func (t tokenservice) VerifyTokenValidity(ctx context.Context, acToken string) (isValid bool) {
	valid := t.tokenRepo.VerifyTokenValidity(ctx, acToken)
	return valid
}
