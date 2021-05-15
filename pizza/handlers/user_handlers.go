package handlers

import (
	"net/http"

	"github.com/VarthanV/pizza/users"
	"github.com/VarthanV/pizza/users/models"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService users.Service
}

func (u UserHandler) SignUpUser(c *gin.Context) {
	var request UserSignupRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": "Please make sure if you have sent all the fields right"})
	}
	// If the request is ok create a user
	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	
	err = u.userService.CreateUser(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
	}
	// If everything went well return a 201 response
	c.Status(http.StatusCreated)
}
