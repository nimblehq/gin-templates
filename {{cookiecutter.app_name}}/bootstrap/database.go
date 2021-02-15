package bootstrap

import (
	"log"
	"strings"

	"github.com/nimblehq/{{cookiecutter.app_name}}/helpers"

	"github.com/gin-gonic/gin"
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
		log.Println(strings.Title(gin.Mode()) + " database connected successfully.")
	}
}
