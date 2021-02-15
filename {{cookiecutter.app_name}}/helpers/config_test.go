package helpers_test

import (
	"github.com/gin-gonic/gin"
	"github.com/nimblehq/xxxx/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	Describe("#GetConfigPrefix", func() {
		Context("given RELEASE gin mode", func() {
			It("returns correct prefix", func() {
				gin.SetMode(gin.ReleaseMode)
				defer gin.SetMode(gin.TestMode)

				result := helpers.GetConfigPrefix()

				Expect(result).To(BeEmpty())
			})
		})

		Context("given DEBUG gin mode", func() {
			It("returns correct prefix", func() {
				gin.SetMode(gin.DebugMode)
				defer gin.SetMode(gin.TestMode)

				result := helpers.GetConfigPrefix()

				Expect(result).To(Equal("debug."))
			})
		})

		Context("given TEST gin mode", func() {
			It("returns correct prefix", func() {
				gin.SetMode(gin.TestMode)
				defer gin.SetMode(gin.TestMode)

				result := helpers.GetConfigPrefix()

				Expect(result).To(Equal("test."))
			})
		})
	})

	// SetupTestEnvironment loaded the config already.
	Describe("#ReadConfig", func() {
		Context("given EXISTING config key", func() {
			It("returns correct config value", func() {
				result := helpers.ReadConfig("database_url")

				Expect(result).To(Equal("postgres://postgres:postgres@0.0.0.0:5433/xxxx_test?sslmode=disable"))
			})
		})

		Context("given NON-EXISTING config key", func() {
			It("returns empty config value", func() {
				result := helpers.ReadConfig("invalid")

				Expect(result).To(BeEmpty())
			})
		})
	})
})
