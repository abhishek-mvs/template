package restaurantService

import (
	"template/internal/app/dto"
	"template/internal/pkg/db"

	"github.com/google/uuid"
)

type RestaurantService struct {
	DB *db.DB
}

func NewRestaurantService(db *db.DB) *RestaurantService {
	return &RestaurantService{
		DB: db,
	}
}

func (r *RestaurantService) GetAllRestaurants() (dto.RestaurantResponse, error) {
	allRestaurants := r.DB.GetAll("restaurants")
	
	restaurants := make([]dto.Restaurant, 0, len(allRestaurants))
	for _, restaurantVal := range allRestaurants {
		restaurant, ok := restaurantVal.(*dto.Restaurant)
		if !ok {
			continue
		}
		restaurants = append(restaurants, *restaurant)
	}

	return dto.RestaurantResponse{
		Restaurants: restaurants,
	}, nil
}

func (r *RestaurantService) GetRestaurantById(id uuid.UUID) (*dto.Restaurant, error) {
	restaurantVal, exists := r.DB.Get("restaurants", id)
	if !exists {
		return nil, nil
	}

	restaurant, ok := restaurantVal.(*dto.Restaurant)
	if !ok {
		return nil, nil
	}

	return restaurant, nil
}

