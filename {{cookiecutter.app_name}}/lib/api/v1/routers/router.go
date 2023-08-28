package routers

import (
	"github.com/nimblehq/{{cookiecutter.app_name}}/lib/api/v1/controllers"

	"github.com/gin-gonic/gin"
)

func CombineRoutes(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")

	v1.GET("/health", controllers.HealthController{}.HealthStatus)
}
