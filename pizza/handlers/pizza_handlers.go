package handlers

import (
	"net/http"

	"github.com/VarthanV/pizza/pizza"
	"github.com/gin-gonic/gin"
)


type PizzaHandler struct {
	pizzaService pizza.Service
}

func NewPizzaHandler(service pizza.Service) *PizzaHandler {
	return &PizzaHandler{
		pizzaService: service,
	}
}


func (p PizzaHandler) GetAllPizzas(c*gin.Context){
	pizzas,err := p.pizzaService.GetAllPizzas(c,0)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"status":"error","error":err.Error()})
		return
	}
	if pizzas == nil {
		c.JSON(http.StatusOK,gin.H{"status":"ok","pizzas":[]string{}})
		return
	}
	c.JSON(http.StatusOK,gin.H{"status":"ok","pizzas":pizzas})
}