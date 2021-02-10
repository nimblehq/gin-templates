package main

import (
	"log"

	"github.com/nimblehq/xxx/config"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	config.ComebineRouter(app)

	err := app.Run()
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
