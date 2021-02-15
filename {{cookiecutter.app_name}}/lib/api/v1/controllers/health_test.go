package controllers_test

import (
	"net/http"

	"github.com/nimblehq/xxxx/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HealthController", func() {
	Describe("GET /health", func() {
		It("returns status OK", func() {
			resp := test.MakeRequest("GET", "/api/v1/health", nil)

			Expect(resp.Code).To(Equal(http.StatusOK))
		})

		It("returns correct response body", func() {
			resp := test.MakeRequest("GET", "/api/v1/health", nil)

			Expect(resp.Body.String()).To(Equal("{\"status\":\"alive\"}"))
		})
	})
})
