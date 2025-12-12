package dishService

import (
	"template/internal/app/dto"
	"template/internal/pkg/db"

	"github.com/google/uuid"
)

type DishService struct {
	DB *db.DB
}

func NewDishService(db *db.DB) *DishService {
	return &DishService{
		DB: db,
	}
}

func (d *DishService) GetDishesByRestaurantId(restaurantId uuid.UUID) (dto.DishResponse, error) {
	allDishes := d.DB.GetAll("dishes")
	
	dishes := make([]dto.Dish, 0)
	for _, dishVal := range allDishes {
		dish, ok := dishVal.(*dto.Dish)
		if !ok {
			continue
		}
		if dish.RestaurantId == restaurantId {
			dishes = append(dishes, *dish)
		}
	}

	return dto.DishResponse{
		Dishes: dishes,
	}, nil
}

func (d *DishService) GetDishById(dishId uuid.UUID) (*dto.Dish, error) {
	dishVal, exists := d.DB.Get("dishes", dishId)
	if !exists {
		return nil, nil
	}

	dish, ok := dishVal.(*dto.Dish)
	if !ok {
		return nil, nil
	}

	return dish, nil
}

