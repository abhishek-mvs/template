package dto

import "github.com/google/uuid"

type HealthCounter struct {
	Id uuid.UUID `json:"id"`
	Counter int `json:"counter"`
}

type HealthResponse struct {
	Id uuid.UUID `json:"id"`
	Counter int `json:"counter"`
	Success bool `json:"success"`
}