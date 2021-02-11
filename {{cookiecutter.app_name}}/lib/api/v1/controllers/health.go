// SAFETODELETE: This file is an example API and can be deleted.
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
