package bootstrap

import (
	apiv1router "github.com/nimblehq/{{cookiecutter.app_name}}/lib/api/v1/routers"
	{% if cookiecutter._web_variant == "y" %}webrouter "github.com/nimblehq/{{cookiecutter.app_name}}/lib/web/routers"{% endif %}

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	apiv1router.ComebineRoutes(r)
	webrouter.ComebineRoutes(r)

	return r
}
