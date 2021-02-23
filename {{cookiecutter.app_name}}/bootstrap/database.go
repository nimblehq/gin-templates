package bootstrap

import (
	"fmt"
	"log"
	"strings"

	"github.com/nimblehq/xxxx/helpers"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() {
	db, err := gorm.Open(postgres.Open(getDatabaseURL()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to %v database: %v", gin.Mode(), err)
	} else {
		viper.Set("database", db)
		log.Println(strings.Title(gin.Mode()) + " database connected successfully.")
	}
}

func getDatabaseURL() string {
	if gin.Mode() == gin.ReleaseMode {
		return viper.GetString("DATABASE_URL")
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		viper.GetString("db_username"),
		viper.GetString("db_password"),
		viper.GetString("db_host"),
		helpers.GetStringConfig("db_port"),
		helpers.GetStringConfig("db_name"),
	)
}
