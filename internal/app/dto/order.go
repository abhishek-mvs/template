package dto

import "github.com/google/uuid"

type Order struct {
	Id           uuid.UUID `json:"id"`
	UserId       string    `json:"user_id"`
	RestaurantId uuid.UUID `json:"restaurant_id"`
	DishId       uuid.UUID `json:"dish_id"`
	DishName     string    `json:"dish_name"`
	Price        float64   `json:"price"`
	Status       string    `json:"status"`
	CreatedAt    string    `json:"created_at"`
}

type CreateOrderRequest struct {
	UserId       string    `json:"user_id" binding:"required"`
	RestaurantId uuid.UUID `json:"restaurant_id" binding:"required"`
	DishId       uuid.UUID `json:"dish_id" binding:"required"`
}

type CreateOrderResponse struct {
	Order Order `json:"order"`
}

type UserOrdersResponse struct {
	Orders []Order `json:"orders"`
}

