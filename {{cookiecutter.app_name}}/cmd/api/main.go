package main

import (
	"log"

	"github.com/nimblehq/{{cookiecutter.app_name}}/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	bootstrap.SetupRouter(app)

	err := app.Run()
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
