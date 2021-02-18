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

func GetBoolConfig(key string) bool {
	return viper.GetBool(GetConfigPrefix() + key)
}

func GetFloatConfig(key string) float64 {
	return viper.GetFloat64(GetConfigPrefix() + key)
}

func GetIntConfig(key string) int {
	return viper.GetInt(GetConfigPrefix() + key)
}

func GetStringConfig(key string) string {
	return viper.GetString(GetConfigPrefix() + key)
}
