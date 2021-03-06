package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
	BaseController
}

func (HealthController) HealthStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "alive",
	})
}
