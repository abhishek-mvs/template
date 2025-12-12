package template

import (
	"template/internal/app/clients/healthClient"
	"template/internal/app/service/healthService"
	"template/internal/app/manager/healthManager"
	"template/internal/app/service/restaurantService"
	"template/internal/app/manager/restaurantManager"
	"template/internal/app/service/dishService"
	"template/internal/app/manager/dishManager"
	"template/internal/app/service/orderService"
	"template/internal/app/manager/orderManager"
	"template/internal/app/controller"
	"template/internal/pkg/db"
)

type Container struct {
	HealthController    *controller.HealthController
	RestaurantController *controller.RestaurantController
	DishController      *controller.DishController
	OrderController     *controller.OrderController
}

func NewContainer() *Container {
	database := db.NewDB()
	redis := db.NewRedis()

	// Seed sample data
	db.SeedData(database)

	// Health check dependencies
	healthClient := healthClient.NewHealthClient()
	healthService := healthService.NewHealthService(healthClient, database, redis)
	healthManager := healthManager.NewHealthManager(healthService)
	healthController := controller.NewHealthController(healthManager)

	// Restaurant dependencies
	restaurantService := restaurantService.NewRestaurantService(database)
	restaurantManager := restaurantManager.NewRestaurantManager(restaurantService)
	restaurantController := controller.NewRestaurantController(restaurantManager)

	// Dish dependencies
	dishService := dishService.NewDishService(database)
	dishManager := dishManager.NewDishManager(dishService)
	dishController := controller.NewDishController(dishManager)

	// Order dependencies
	orderService := orderService.NewOrderService(database, dishService, restaurantService)
	orderManager := orderManager.NewOrderManager(orderService)
	orderController := controller.NewOrderController(orderManager)

	return &Container{
		HealthController:    healthController,
		RestaurantController: restaurantController,
		DishController:      dishController,
		OrderController:     orderController,
	}
}