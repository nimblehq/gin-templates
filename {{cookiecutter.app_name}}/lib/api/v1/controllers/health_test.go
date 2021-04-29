package controllers_test

import (
	"net/http"

	"github.com/nimblehq/{{cookiecutter.app_name}}/lib/api/v1/controllers"
	"github.com/nimblehq/{{cookiecutter.app_name}}/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HealthController", func() {
	Describe("GET /health", func() {
		It("returns status OK", func() {
			c, resp := test.CreateGinTestContext()
			healthController := controllers.HealthController{}

			healthController.HealthStatus(c)

			Expect(resp.Code).To(Equal(http.StatusOK))
		})

		It("returns correct response body", func() {
			c, resp := test.CreateGinTestContext()
			healthController := controllers.HealthController{}

			healthController.HealthStatus(c)

			Expect(resp.Body.String()).To(Equal("{\"status\":\"alive\"}"))
		})
	})
})
