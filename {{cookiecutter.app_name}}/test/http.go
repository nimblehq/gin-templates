package test

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func CreateGinTestContext() (*gin.Context, *httptest.ResponseRecorder) {
	resp := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(resp)

	return c, resp
}
