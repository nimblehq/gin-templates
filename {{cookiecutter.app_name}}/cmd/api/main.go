package main

import (
	"fmt"
	{% if cookiecutter.use_logrus == "no" %}"log"{% endif %}

	"github.com/nimblehq/{{cookiecutter.app_name}}/bootstrap"
	{% if cookiecutter.use_logrus == "yes" %}"github.com/nimblehq/{{cookiecutter.app_name}}/helpers/log"{% endif %}

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	bootstrap.LoadConfig()

	bootstrap.InitDatabase()

	r := bootstrap.SetupRouter()

	err := r.Run(getAppPort())
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

func getAppPort() string {
	if gin.Mode() == gin.ReleaseMode {
		return fmt.Sprint(":", viper.GetString("PORT"))
	}

	return fmt.Sprint(":", viper.GetString("app_port"))
}
