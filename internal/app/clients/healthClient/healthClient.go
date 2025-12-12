package healthClient

type HealthClient struct {
}

func NewHealthClient() *HealthClient {
	return &HealthClient{}
}

func (h *HealthClient) HealthCheck() (bool, error) {
	return true, nil
}