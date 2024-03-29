package bootstrap

import (
	{% if cookiecutter.use_logrus == "no" %}"log"
	{% endif %}
	{% if cookiecutter.use_logrus == "yes" %}"github.com/nimblehq/{{cookiecutter.app_name}}/helpers/log"
	{% endif %}
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("app")
	viper.SetConfigType("toml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}
}
