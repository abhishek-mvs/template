package healthService

import "template/internal/app/clients/healthClient"

type HealthService struct {
	HealthClient *healthClient.HealthClient
}

func NewHealthService(healthClient *healthClient.HealthClient) *HealthService {
	return &HealthService{
		HealthClient: healthClient,
	}
}

func (h *HealthService) HealthCheck() (bool, error) {
	return h.HealthClient.HealthCheck()
}