package main

import (
	"log"

	"github.com/nimblehq/xxxx/bootstrap"
)

func main() {
	bootstrap.LoadConfig()

	r := bootstrap.SetupRouter()

	err := r.Run()
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
