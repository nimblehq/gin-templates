package helpers

import (
	"errors"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// Get database connection from viper config
func GetDBConnection() (*gorm.DB, error) {
	dbConnection, ok := viper.Get("database").(*gorm.DB)
	if !ok {
		return nil, errors.New("Failed to get database connection: connection is not type of *gorm.DB")
	}

	return dbConnection, nil
}
