package test

import (
	"github.com/nimblehq/xxxx/bootstrap"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func SetupTestEnvironment() {
	gin.SetMode(gin.TestMode)

	Router = bootstrap.SetupRouter()
}
