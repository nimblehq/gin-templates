package bootstrap

import (
	"log"

	apiv1router "github.com/nimblehq/xxxx/lib/api/v1/routers"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	viper.SetConfigName("config")
	viper.SetConfigType("conf")
	viper.AddConfigPath("./config")
	// viper.AutomaticEnv()
	_ = viper.ReadInConfig()
	log.Printf("=============== %v", gin.Mode())
	log.Printf("=============== %v", viper.GetString(gin.Mode()+".dbURL"))

	r := gin.Default()

	apiv1router.ComebineRoutes(r)

	return r
}
