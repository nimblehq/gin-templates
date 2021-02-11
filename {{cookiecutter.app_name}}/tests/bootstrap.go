package tests

import (
	"github.com/nimblehq/{{cookiecutter.app_name}}/bootstrap"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func SetupTestEnvironment() {
	gin.SetMode(gin.TestMode)

	Router = bootstrap.SetupRouter()
}
