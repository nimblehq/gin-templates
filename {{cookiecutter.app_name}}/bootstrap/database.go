package bootstrap

import (
	"log"

	"github.com/nimblehq/xxxx/helpers"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Init database and set to viper config
func InitDatabase() {
	db, err := gorm.Open(postgres.Open(helpers.GetStringConfig("database_url")), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	} else {
		viper.Set("database", db)
		log.Println("Database connected successfully.")
	}
}
