package utils

import (
	"github.com/VarthanV/pizza/users/models"
	"github.com/gin-gonic/gin"
)

type UtilityService interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
	CreateToken(user models.User) (accessToken string, refreshToken string, atExpiresAt int64, rtExpiresAt int64, err error)
	GetUserFromContext(c *gin.Context) (userID string)
}
