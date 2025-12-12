package template

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) *gin.RouterGroup {
	container := NewContainer()
	v1 := router.Group("/v1")
	externalRoutes(v1, container)
	return v1
}

func externalRoutes(routerGroup *gin.RouterGroup, container *Container) {
	// Health check
	routerGroup.GET("/health", container.HealthController.HealthCheck)

	// Restaurant routes
	routerGroup.GET("/restaurants", container.RestaurantController.GetAllRestaurants)

	// Dish routes
	routerGroup.GET("/restaurants/:id/dishes", container.DishController.GetDishesByRestaurantId)

	// Order routes
	routerGroup.POST("/orders", container.OrderController.CreateOrder)
	routerGroup.GET("/users/:userId/orders", container.OrderController.GetUserOrders)
}
