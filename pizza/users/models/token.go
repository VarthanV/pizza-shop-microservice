package models

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}


type TokenRepository interface {
	CreateToken(user User) (tokenDetails TokenDetails,err error)
	VerifyTokenValidity(token TokenDetails) (isValid bool)
}