package controllers

import (
	"github.com/gin-gonic/gin"
)

type BaseController struct{}

func (BaseController) Data(c *gin.Context, data gin.H) gin.H {
	data["CurrentPath"] = c.Request.URL

	return data
}
