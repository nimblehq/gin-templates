package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func GetConfigPrefix() string {
	if gin.Mode() == "release" {
		return ""
	}

	return gin.Mode() + "."
}

func ReadConfig(key string) string {
	return viper.GetString(GetConfigPrefix() + key)
}
