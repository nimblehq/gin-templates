package bootstrap

import (
	apiv1router "github.com/nimblehq/{{cookiecutter.app_name}}/lib/api/v1/routers"
	{% if cookiecutter._web_variant == "yes" %}webrouter "github.com/nimblehq/{{cookiecutter.app_name}}/lib/web/routers"
	{% endif %}
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	apiv1router.ComebineRoutes(r)
	{% if cookiecutter._web_variant == "yes" %}webrouter.ComebineRoutes(r)
	{% endif %}
	return r
}
