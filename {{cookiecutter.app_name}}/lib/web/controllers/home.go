package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct {
	BaseController
}

func (hc HomeController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index", hc.Data(c, gin.H{
		"BodyClass": "home index",
	}))
}
