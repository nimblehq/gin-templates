package bootstrap

import (
	"github.com/nimblehq/xxx/config"

	"github.com/gin-gonic/gin"
)

func SetupRouter(app *gin.Engine) {
	config.ComebineRouter(app)
}
