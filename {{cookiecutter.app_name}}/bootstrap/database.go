package bootstrap

import (
	"github.com/nimblehq/{{cookiecutter.app_name}}/database"
)

func InitDatabase(databaseURL string) {
	database.InitDatabase(databaseURL)
}
