package test

import (
	"os"
	"path/filepath"
	"runtime"
	{% if cookiecutter.use_logrus != "y" %}"log"{% endif %}

	"github.com/nimblehq/{{cookiecutter.app_name}}/bootstrap"
	{% if cookiecutter.use_logrus == "y" %}"github.com/nimblehq/{{cookiecutter.app_name}}/helpers/log"{% endif %}

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func SetupTestEnvironment() {
	gin.SetMode(gin.TestMode)

	setRootDir()

	bootstrap.LoadConfig()

	bootstrap.InitDatabase()
}

func setRootDir() {
	_, currentFile, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(currentFile), "../")

	err := os.Chdir(root)
	if err != nil {
		log.Fatal("Failed to set root directory: ", err)
	}
}
