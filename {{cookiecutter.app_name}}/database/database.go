package database

import (
	"fmt"
	"strings"
	{% if cookiecutter.use_logrus == "no" %}"log"
	{% endif %}

	"github.com/nimblehq/{{cookiecutter.app_name}}/helpers"
	{% if cookiecutter.use_logrus == "yes" %}"github.com/nimblehq/{{cookiecutter.app_name}}/helpers/log"
	{% endif %}

	"github.com/gin-gonic/gin"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func InitDatabase(databaseURL string) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to %v database: %v", gin.Mode(), err)
	} else {
		viper.Set("database", db)
		log.Println(strings.Title(gin.Mode()) + " database connected successfully.")
	}

	migrateDB(db)
}

func migrateDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to convert gormDB to sqlDB: %v", err)
	}

	err = goose.Up(sqlDB, "database/migrations", goose.WithAllowMissing())
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Migrated database successfully.")
}

func GetDB() *gorm.DB {
	if database == nil {
		InitDatabase(GetDatabaseURL())
	}

	return database
}

func GetDatabaseURL() string {
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
