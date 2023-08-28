package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type OpenAPIController struct{}

func (OpenAPIController) Show(c *gin.Context) {
	c.HTML(http.StatusOK, "show.html", gin.H{})
}
