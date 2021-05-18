package handlers

import (
	"net/http"

	"github.com/VarthanV/pizza/pizza"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

type PizzaHandler struct {
	pizzaService pizza.Service
}

func NewPizzaHandler(service pizza.Service) *PizzaHandler {
	return &PizzaHandler{
		pizzaService: service,
	}
}

func (p PizzaHandler) GetAllPizzas(c *gin.Context) {
	var isVeg int
	isVegetarian := c.DefaultQuery("is_vegeterian", "false")
	if isVegetarian == "true" {
		isVeg = 1
	} else {
		isVeg = 0
	}
	glog.Infof("Is vegeterian: %d ?", isVeg)
	pizzas, err := p.pizzaService.GetAllPizzas(c, isVeg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
		return
	}
	if pizzas == nil {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "pizzas": []string{}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "pizzas": pizzas})
}
