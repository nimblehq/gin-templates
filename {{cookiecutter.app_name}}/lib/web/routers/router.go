package routers

import (
	"html/template"
	"net/url"

	"github.com/nimblehq/{{cookiecutter.app_name}}/lib/web/controllers"

	"github.com/gin-gonic/gin"
	eztemplate "github.com/michelloworld/ez-gin-template"
)

func CombineRoutes(engine *gin.Engine) {
	// Register HTML renderer
	htmlRender := eztemplate.New()
	htmlRender.Debug = gin.IsDebugging()
	htmlRender.TemplatesDir = "lib/web/views/"
	htmlRender.Layout = "layouts/application"
	htmlRender.TemplateFuncMap = template.FuncMap{
		"isActive": isActive,
	}
	engine.HTMLRender = htmlRender.Init()

	// Assets
	router := engine.Group("/")
	router.Static("/static", "./static")
	router.Static("/assets/images", "./lib/web/assets/images")

	// Routes
	router.GET("/", controllers.HomeController{}.Index)
}

func isActive(currentPath *url.URL, path string) bool {
	return currentPath.String() == path
}
