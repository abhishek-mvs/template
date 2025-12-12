package template

import (
	"template/internal/app/clients/healthClient"
	"template/internal/app/service/healthService"
	"template/internal/app/manager/healthManager"
	"template/internal/app/controller"
)


type Container struct {
	HealthController *controller.HealthController
}

func NewContainer() *Container {
	healthClient := healthClient.NewHealthClient()
	healthService := healthService.NewHealthService(healthClient)
	healthManager := healthManager.NewHealthManager(healthService)
	healthController := controller.NewHealthController(healthManager)
	return &Container{
		HealthController: healthController,
	}
}