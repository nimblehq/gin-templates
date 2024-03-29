package bootstrap

import (
	{%- if cookiecutter._api_variant == "yes" -%}apidocsrouter "github.com/nimblehq/{{cookiecutter.app_name}}/lib/api/docs/routers"{%- endif %}
	apiv1router "github.com/nimblehq/{{cookiecutter.app_name}}/lib/api/v1/routers"
	{% if cookiecutter._web_variant == "yes" %}webrouter "github.com/nimblehq/{{cookiecutter.app_name}}/lib/web/routers"
	{% endif %}

	{%- if cookiecutter._api_variant == "yes" -%}"github.com/gin-gonic/contrib/static"{%- endif %}
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	{% if cookiecutter._api_variant == "yes" -%}r.Use(static.Serve("/", static.LocalFile("./public", true))){%- endif %}

	apiv1router.CombineRoutes(r)
	{% if cookiecutter._api_variant == "yes" -%}apidocsrouter.CombineRoutes(r){%- endif %}
	{% if cookiecutter._web_variant == "yes" %}webrouter.CombineRoutes(r)
	{% endif %}
	return r
}
