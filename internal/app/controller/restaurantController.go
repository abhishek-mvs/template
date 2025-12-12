package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"template/internal/app/manager/restaurantManager"
)

type RestaurantController struct {
	restaurantManager *restaurantManager.RestaurantManager
}

func NewRestaurantController(restaurantManager *restaurantManager.RestaurantManager) *RestaurantController {
	return &RestaurantController{
		restaurantManager: restaurantManager,
	}
}

func (r *RestaurantController) GetAllRestaurants(c *gin.Context) {
	response, err := r.restaurantManager.GetAllRestaurants()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

