package dto

import "github.com/google/uuid"

type Restaurant struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Address     string    `json:"address"`
	Rating      float64   `json:"rating"`
}

type RestaurantResponse struct {
	Restaurants []Restaurant `json:"restaurants"`
}

