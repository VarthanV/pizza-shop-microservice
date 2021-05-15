package models

import "context"

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AtExpiresAt int64
	RtExpiresAt int64
}


type TokenRepository interface {
	CreateToken( ctx context.Context ,user User) (tokenDetails TokenDetails,err error)
	VerifyTokenValidity(token TokenDetails) (isValid bool)
}