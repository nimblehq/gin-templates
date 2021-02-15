package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func ReadConfig(key string) string {
	var prefix string

	if gin.Mode() != "release" {
		prefix = gin.Mode() + "."
	}

	return viper.GetString(prefix + key)
}
