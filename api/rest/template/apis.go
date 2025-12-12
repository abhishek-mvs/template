package template

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) *gin.RouterGroup {
	container := NewContainer()
	v1 := router.Group("/v1")
	externalRoutes(v1, container)
	return v1
}

func externalRoutes(routerGroup *gin.RouterGroup, container *Container) {
	routerGroup.GET("/health", container.HealthController.HealthCheck)
}
