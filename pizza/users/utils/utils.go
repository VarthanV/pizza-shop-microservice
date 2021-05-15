package utils

import (
	"time"

	"github.com/VarthanV/pizza/users/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/glog"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateToken(user models.User) (accessToken string, refreshToken string, atExpiresAt int64, rtExpiresAt int64 , err error) {
	fiftyMinutesFromNow := time.Now().Add(time.Minute * 5).Unix()
	fiveDaysFromNow := time.Now().Add(time.Hour * 24 * 5).Unix()

	// AccessToken claims
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = user.ID
	atClaims["email"] = user.Email
	atClaims["exp"] = fiftyMinutesFromNow

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	// Sign it with the secret hash
	acToken, err := at.SignedString([]byte("supersecreeett"))
	if err != nil {
		glog.Fatal("Error while creating access token...", err)
		return "","",0, 0 ,err
	}
	// RefreshToken Claims
	rtClaims := jwt.MapClaims{}
	rtClaims["authorized"] = true
	atClaims["user_id"] = user.ID
	atClaims["email"] = user.Email
	atClaims["exp"] = fiveDaysFromNow
	rtClaims["ref_uuid"] = uuid.New().String()

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	refToken, err := rt.SignedString([]byte("supersecreetrefesh"))
	if err != nil {
		glog.Fatal("Error while creating refresh token...", err)
		return "","",0, 0 ,err
	}
	return acToken, refToken,fiftyMinutesFromNow,fiveDaysFromNow, nil
}
