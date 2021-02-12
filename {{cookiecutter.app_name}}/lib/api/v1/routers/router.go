package routers

import (
	"github.com/nimblehq/{{cookiecutter.app_name}}/lib/api/v1/controllers"

	"github.com/gin-gonic/gin"
)

func ComebineRoutes(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")

	// SAFETODELETE: This is an example API router and can be deleted.
	v1.GET("/health", controllers.HealthController{}.HealthStatus)
}
