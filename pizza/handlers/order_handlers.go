package handlers

import (
	"github.com/VarthanV/pizza/pizza"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService pizza.OrderService
}

func NewOrderHandler(orderService pizza.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}

}

func (o OrderHandler) CreateOrder(c *gin.Context) {

}
