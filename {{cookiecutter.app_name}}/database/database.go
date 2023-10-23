package database

import (
	"fmt"
	{% if cookiecutter.use_logrus == "no" %}"log"
	{% endif %}

	"github.com/nimblehq/{{cookiecutter.app_name}}/helpers"
	{% if cookiecutter.use_logrus == "yes" %}"github.com/nimblehq/{{cookiecutter.app_name}}/helpers/log"
	{% endif %}

	"github.com/gin-gonic/gin"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
	"golang.org/x/text/cases"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

const databaseDir = "database/migrations"

func InitDatabase(databaseURL string) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to %v database: %v", gin.Mode(), err)
	} else {
		viper.Set("database", db)
		caser := cases.Title(language.English)
		log.Println(caser.String(gin.Mode()) + " database connected successfully.")
	}

	migrateDB(db)
}

func migrateDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to convert gormDB to sqlDB: %v", err)
	}

	if migrationFileExist() {
		err = goose.Up(sqlDB, databaseDir, goose.WithAllowMissing())
		if err != nil {
			log.Fatalf("Failed to migrate the database: %v", err)
		}

		log.Println("Database migrated successfully.")
	} else {
		log.Println("NO migration files")
	}
}

func migrationFileExist() bool {
	files, err := os.ReadDir(databaseDir)
	if err != nil {
		log.Fatalf("Missing migration directory: %v", err)
	}

	return len(files) > 0 && files[0].Name() != ".keep"
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
