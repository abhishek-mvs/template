package dishManager

import (
	"template/internal/app/dto"
	"template/internal/app/service/dishService"

	"github.com/google/uuid"
)

type DishManager struct {
	dishService *dishService.DishService
}

func NewDishManager(dishService *dishService.DishService) *DishManager {
	return &DishManager{
		dishService: dishService,
	}
}

func (d *DishManager) GetDishesByRestaurantId(restaurantId uuid.UUID) (dto.DishResponse, error) {
	return d.dishService.GetDishesByRestaurantId(restaurantId)
}
