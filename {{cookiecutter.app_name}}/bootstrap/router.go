package bootstrap

import (
	{%- if cookiecutter.use_openapi == "yes" -%}apidocsrouter "github.com/nimblehq/{{cookiecutter.app_name}}/lib/api/docs/routers"{%- endif %}
	apiv1router "github.com/nimblehq/{{cookiecutter.app_name}}/lib/api/v1/routers"
	{% if cookiecutter._web_variant == "yes" %}webrouter "github.com/nimblehq/{{cookiecutter.app_name}}/lib/web/routers"
	{% endif %}

	{%- if cookiecutter.use_openapi == "yes" -%}"github.com/gin-gonic/contrib/static"{%- endif %}
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	{% if cookiecutter.use_openapi == "yes" -%}r.Use(static.Serve("/", static.LocalFile("./public", true))){%- endif %}

	apiv1router.ComebineRoutes(r)
	{% if cookiecutter.use_openapi == "yes" -%}apidocsrouter.ComebineRoutes(r){%- endif %}
	{% if cookiecutter._web_variant == "yes" %}webrouter.ComebineRoutes(r)
	{% endif %}
	return r
}
