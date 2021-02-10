package config

import (
	"github.com/nimblehq/xxx/lib/api/v1/controllers"

	"github.com/gin-gonic/gin"
)

func ComebineRouter(engine *gin.Engine) {
	engine.GET("/health", controllers.HealthController{}.HealthStatus)
}
