package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"

	"github.com/VarthanV/pizza/shared"
	"github.com/VarthanV/pizza/users/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/glog"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type utilityservice struct {
	constants *shared.SharedConstants
}

func NewUtilityService(constants *shared.SharedConstants) UtilityService {
	return &utilityservice{
		constants: constants,
	}
}

func (u utilityservice) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (u utilityservice) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (u utilityservice) CreateToken(user models.User) (accessToken string, refreshToken string, atExpiresAt int64, rtExpiresAt int64, err error) {
	fiveHoursFromNow := time.Now().Add(time.Hour * 5).Unix()
	fiveDaysFromNow := time.Now().Add(time.Hour * 24 * 5).Unix()

	// AccessToken claims
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = user.ID
	atClaims["email"] = user.Email
	atClaims["exp"] = fiveHoursFromNow

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	// Sign it with the secret hash
	acToken, err := at.SignedString([]byte(u.constants.AccessTokenSecretKey))
	if err != nil {
		glog.Fatal("Error while creating access token...", err)
		return "", "", 0, 0, err
	}
	// RefreshToken Claims
	rtClaims := jwt.MapClaims{}
	rtClaims["authorized"] = true
	atClaims["user_id"] = user.ID
	atClaims["email"] = user.Email
	atClaims["exp"] = fiveDaysFromNow
	rtClaims["ref_uuid"] = uuid.New().String()

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	refToken, err := rt.SignedString([]byte(u.constants.RefreshTokenSecretKey))
	if err != nil {
		glog.Fatal("Error while creating refresh token...", err)
		return "", "", 0, 0, err
	}
	return acToken, refToken, fiveHoursFromNow, fiveDaysFromNow, nil
}

func (u utilityservice) GetUserFromContext(c *gin.Context) (userID string) {
	user, _ := c.Get("userID")
	ID := fmt.Sprintf("%v", user)
	return ID
}
