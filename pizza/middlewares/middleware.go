package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/VarthanV/pizza/shared"
	"github.com/VarthanV/pizza/users"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

type Middleware struct {
	constants    *shared.SharedConstants
	tokenService users.TokenService
}

func NewMiddleware(constants *shared.SharedConstants, tokenService users.TokenService) Service {
	return &Middleware{
		constants:    constants,
		tokenService: tokenService,
	}
}

func (m Middleware) VerifyTokenMiddleware(c *gin.Context) {
	var tokenString string
	//var claims jwt.Claims

	authorizationHeader := c.Request.Header.Get("Authorization")
	if authorizationHeader == "" {
		glog.Info("Blank Authorization header")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	glog.Info("Authorization is...", authorizationHeader)

	bearToken := c.Request.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		tokenString = strArr[1]
	}
	// First look up in the store if the token is not expired or something

	isValid := m.tokenService.VerifyTokenValidity(c, tokenString)
	if !isValid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		if _, ok := t.Claims.(jwt.Claims); !ok && !t.Valid {
			return nil, nil
		}
		claims, ok := t.Claims.(jwt.MapClaims)
		if ok {
			c.Set("userID", claims["user_id"])
		}
		c.Next()
		return []byte(m.constants.AccessTokenSecretKey), nil
	})
	if err != nil {
		glog.Errorf("Error unable to parse token %s", err.Error())
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	glog.Info("Token is...", token)
}
