package bootstrap

import (
	"log"

	"github.com/nimblehq/xxxx/helpers"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique;notNull;index"`
	Password string `gorm:"notNull"`
}

// Init database and set to viper config
func InitDatabase() {
	db, err := gorm.Open(postgres.Open(helpers.ReadConfig("database_url")), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	} else {
		log.Println("Database connected successfully.")
	}

	viper.Set("database", db)
}
