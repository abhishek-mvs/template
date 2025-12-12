package healthManager

import "template/internal/app/service/healthService"

type HealthManager struct {
	healthService *healthService.HealthService
}

func NewHealthManager(healthService *healthService.HealthService) *HealthManager {
	return &HealthManager{
		healthService: healthService,
	}
}

func (h *HealthManager) HealthCheck() (bool, error) {
	return h.healthService.HealthCheck()
}