package controller

import (
	"net/http"

	"template/internal/app/dto"
	"template/internal/app/manager/orderManager"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderManager *orderManager.OrderManager
}

func NewOrderController(orderManager *orderManager.OrderManager) *OrderController {
	return &OrderController{
		orderManager: orderManager,
	}
}

func (o *OrderController) CreateOrder(c *gin.Context) {
	var req dto.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := o.orderManager.CreateOrder(req.UserId, req.RestaurantId, req.DishId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, response)
}

func (o *OrderController) GetUserOrders(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user id is required"})
		return
	}

	response, err := o.orderManager.GetUserOrders(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
