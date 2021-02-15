package helpers

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// Get database connection from viper config
func GetDBConnection() *gorm.DB {
	return viper.Get("database").(*gorm.DB)
}
