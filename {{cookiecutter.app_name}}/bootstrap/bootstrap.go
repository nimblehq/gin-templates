package bootstrap

import (
	"github.com/nimblehq/{{cookiecutter.app_name}}/database"
)

func Init() {
	LoadConfig()

	InitDatabase(database.GetDatabaseURL())
}
