package bootstrap

import (
	apiv1router "github.com/nimblehq/{{cookiecutter.app_name}}/lib/api/v1/routers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(engine *gin.Engine) {
	apiv1router.ComebineRoutes(engine)
}
