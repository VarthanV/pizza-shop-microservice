package handlers

import (
	"github.com/VarthanV/pizza/pizza"
	"github.com/VarthanV/pizza/pizza/services"
	"github.com/VarthanV/pizza/users/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderHandler struct {
	orderService     pizza.OrderService
	orderItemService services.OrderItemService
	utilityService   utils.UtilityService
}

func NewOrderHandler(orderService pizza.OrderService, orderItemService services.OrderItemService, utilityService utils.UtilityService) *OrderHandler {
	return &OrderHandler{
		orderService:     orderService,
		orderItemService: orderItemService,
		utilityService:   utilityService,
	}

}

func (o OrderHandler) CreateOrder(c *gin.Context) {
	userID := o.utilityService.GetUserFromContext(c)
	if userID == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	err := o.orderService.CreateOrder(c, userID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}
