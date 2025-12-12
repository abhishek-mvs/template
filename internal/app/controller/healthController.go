package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"template/internal/app/manager/healthManager"
)

type HealthController struct {
	healthManager *healthManager.HealthManager
}

func NewHealthController(healthManager *healthManager.HealthManager) *HealthController {
	return &HealthController{
		healthManager: healthManager,
	}
}

func (h *HealthController) HealthCheck(c *gin.Context) {
	ok, err := h.healthManager.HealthCheck()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !ok {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Service is not available"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}