package middlewares

import "github.com/gin-gonic/gin"

type Service interface {
	VerifyTokenMiddleware(c *gin.Context)
}
