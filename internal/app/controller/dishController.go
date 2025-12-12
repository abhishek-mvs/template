package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"template/internal/app/manager/dishManager"
	"github.com/google/uuid"
)

type DishController struct {
	dishManager *dishManager.DishManager
}

func NewDishController(dishManager *dishManager.DishManager) *DishController {
	return &DishController{
		dishManager: dishManager,
	}
}

func (d *DishController) GetDishesByRestaurantId(c *gin.Context) {
	restaurantIdStr := c.Param("id")
	restaurantId, err := uuid.Parse(restaurantIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant id"})
		return
	}

	response, err := d.dishManager.GetDishesByRestaurantId(restaurantId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

