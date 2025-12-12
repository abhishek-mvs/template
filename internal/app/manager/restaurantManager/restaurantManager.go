package restaurantManager

import (
	"template/internal/app/dto"
	"template/internal/app/service/restaurantService"
)

type RestaurantManager struct {
	restaurantService *restaurantService.RestaurantService
}

func NewRestaurantManager(restaurantService *restaurantService.RestaurantService) *RestaurantManager {
	return &RestaurantManager{
		restaurantService: restaurantService,
	}
}

func (r *RestaurantManager) GetAllRestaurants() (dto.RestaurantResponse, error) {
	return r.restaurantService.GetAllRestaurants()
}

