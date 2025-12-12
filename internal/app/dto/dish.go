package dto

import "github.com/google/uuid"

type Dish struct {
	Id           uuid.UUID `json:"id"`
	RestaurantId uuid.UUID `json:"restaurant_id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Price        float64   `json:"price"`
	IsAvailable  bool      `json:"is_available"`
}

type DishResponse struct {
	Dishes []Dish `json:"dishes"`
}

