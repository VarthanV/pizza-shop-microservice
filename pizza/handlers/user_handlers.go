package handlers

import (
	"net/http"

	"github.com/VarthanV/pizza/users"
	"github.com/VarthanV/pizza/users/models"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

type UserHandler struct {
	userService users.Service
}

// Initialize the handler with user service
func NewUserHandler(service users.Service) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (u UserHandler) SignUpUser(c *gin.Context) {
	var request UserSignupRequest
	err := c.BindJSON(&request)
	if err != nil {
		glog.Info("Failed binding json...", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": "Please make sure if you have sent all the fields right"})
		return
	}
	// If the request is ok create a user
	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		PhoneNumber: request.PhoneNumber,
	}

	err = u.userService.CreateUser(c, user)
	if err != nil {
		if err.Error() == "conflict" {
			c.JSON(http.StatusConflict, gin.H{"status": "error", "error": "Please login with your account"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
		return
	}
	// If everything went well return a 201 response
	c.Status(http.StatusCreated)
}

func (u UserHandler) Test(c *gin.Context) {
	c.Status(http.StatusOK)
}
