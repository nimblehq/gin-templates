module github.com/nimblehq/{{cookiecutter.app_name}}

go 1.20

require (
	github.com/gin-gonic/gin v1.8.1
	github.com/onsi/ginkgo v1.15.0
	github.com/onsi/gomega v1.10.5
	github.com/pelletier/go-toml v1.8.1 // indirect
	github.com/spf13/viper v1.7.1
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.20.12
	{% if cookiecutter.use_logrus == "yes" -%}github.com/sirupsen/logrus v1.8.1{%- endif %}
)
