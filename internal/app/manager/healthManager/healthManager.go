package healthManager

import (
	"template/internal/app/dto"
	"template/internal/app/service/healthService"
)

type HealthManager struct {
	healthService *healthService.HealthService
}

func NewHealthManager(healthService *healthService.HealthService) *HealthManager {
	return &HealthManager{
		healthService: healthService,
	}
}

func (h *HealthManager) HealthCheck() (dto.HealthResponse, error) {
	return h.healthService.HealthCheck()
}