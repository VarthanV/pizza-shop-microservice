package handlers

import (
	"fmt"
	"github.com/VarthanV/pizza/pizza/services"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
)

type CartHandler struct {
	cartService services.CartService
}

func NewCartHandler(svc services.CartService) *CartHandler {
	return &CartHandler{
		cartService: svc,
	}
}

func (cart CartHandler) AddToCart(c *gin.Context) {
	var request AddToCartRequest
	err := c.BindJSON(&request)
	if err != nil {
		glog.Errorf("Error binding json.. %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	user, _ := c.Get("userID")
	userID := fmt.Sprintf("%v", user)
	glog.Infof("User from context is %s", user)
	//	Validate the request body and pass to the cart service
	err = request.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}
	err = cart.cartService.AddItem(c, request.PizzaID, userID, request.Quantity, request.Price)
	if err != nil {
		if err.Error() == "item-conflict" {
			c.JSON(http.StatusConflict, gin.H{"status": "err", "error": "Item already exists"})
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
	return
}
