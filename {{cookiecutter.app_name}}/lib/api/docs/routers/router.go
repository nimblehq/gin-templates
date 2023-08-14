package routers

import (
	"github.com/nimblehq/{{cookiecutter.app_name}}/lib/api/docs/controllers"

	"github.com/gin-gonic/gin"
)

func ComebineRoutes(engine *gin.Engine) {
	docs := engine.Group("/api/docs")

	engine.LoadHTMLFiles("lib/api/docs/views/openapi/show.html")
	docs.GET("/openapi", controllers.OpenAPIController{}.Show)
}
