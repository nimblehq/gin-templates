package test

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/nimblehq/{{cookiecutter.app_name}}/bootstrap"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func SetupTestEnvironment() {
	gin.SetMode(gin.TestMode)

	setRootDir()

	bootstrap.LoadConfig()

	bootstrap.InitDatabase()

	Router = bootstrap.SetupRouter()
}

func setRootDir() {
	_, currentFile, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(currentFile), "../")

	err := os.Chdir(root)
	if err != nil {
		log.Fatal("Failed to set root directory: ", err)
	}
}
