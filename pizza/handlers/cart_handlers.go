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

func (cart CartHandler) GetCart(c *gin.Context) {
	user, _ := c.Get("userID")
	userID := fmt.Sprintf("%v", user)
	cartResult, err := cart.cartService.GetCart(c, userID)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "ok", "cart": []string{}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "cart": cartResult})
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

func (cart CartHandler) EditCart(c *gin.Context) {
	var request EditCartRequest
	err := c.BindJSON(&request)
	if err != nil {
		glog.Errorf("Unable to bind %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "eror": err.Error()})
		return
	}
	user, _ := c.Get("userID")
	userID := fmt.Sprintf("%v", user)
	err = request.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}
	err = cart.cartService.EditItem(c, request.ID, request.PizzaID, request.Quantity, request.Price, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
	return
}
