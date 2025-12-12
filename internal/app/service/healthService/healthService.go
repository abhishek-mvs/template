package healthService

import (
	"errors"
	"template/internal/app/clients/healthClient"
	"template/internal/app/dto"
	"template/internal/pkg/db"

	"github.com/google/uuid"
)

type HealthService struct {
	HealthClient *healthClient.HealthClient
	DB           *db.DB
	Redis        *db.Redis
}

func NewHealthService(healthClient *healthClient.HealthClient, db *db.DB, redis *db.Redis) *HealthService {
	return &HealthService{
		HealthClient: healthClient,
		DB:           db,
		Redis:        redis,
	}
}

func (h *HealthService) HealthCheck() (dto.HealthResponse, error) {
	healthIdVal, exists := h.Redis.Get("health_counter_id")
	var healthId uuid.UUID

	if !exists {
		healthCounter := &dto.HealthCounter{
			Counter: 0,
		}
		id, err := h.DB.Save("health_counter", healthCounter)
		if err != nil {
			return dto.HealthResponse{
				Counter: 0,
				Success: false,
			}, err
		}
		healthCounter.Id = id
	
		h.Redis.Set("health_counter_id", id)
		healthId = id
	} else {
		
		var ok bool
		healthId, ok = healthIdVal.(uuid.UUID)
		if !ok {
			return dto.HealthResponse{
				Counter: 0,
				Success: false,
			}, errors.New("health counter id is not a valid uuid")
		}
	}

	healthCounterVal, exists := h.DB.Get("health_counter", healthId)
	if !exists {
		return dto.HealthResponse{
			Counter: 0,
			Success: false,
		}, errors.New("health counter not found")
	}

	healthCounter, ok := healthCounterVal.(*dto.HealthCounter)
	if !ok {
		return dto.HealthResponse{
			Counter: 0,
			Success: false,
		}, errors.New("health counter is not a valid dto.HealthCounter")
	}

	healthCounter.Counter++
	err := h.DB.Update("health_counter", healthId, healthCounter)
	if err != nil {
		return dto.HealthResponse{
			Counter: 0,
			Success: false,
		}, err
	}

	success, err := h.HealthClient.HealthCheck()
	if err != nil {
		return dto.HealthResponse{
			Counter: healthCounter.Counter,
			Success: false,
		}, err
	}
	return dto.HealthResponse{
		Counter: healthCounter.Counter,
		Success: success,
	}, nil
}
