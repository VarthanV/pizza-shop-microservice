package middlewares

import (
	"net/http"

	"github.com/VarthanV/pizza/shared"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

type Middleware struct {

	constants  *shared.SharedConstants

}

func NewMiddleware (constants *shared.SharedConstants) Service{
	return &Middleware{
		constants: constants,
	}
}


func (m Middleware) VerifyTokenMiddleware(c*gin.Context){
	authorizationHeader := c.Request.Header.Get("Authorization")
	if authorizationHeader == "" {
		glog.Info("Blank Authorization header")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	glog.Info("Authorization is...",authorizationHeader)
	c.Next()

}