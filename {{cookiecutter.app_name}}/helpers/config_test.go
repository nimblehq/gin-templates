package helpers_test

import (
	"github.com/nimblehq/test-gin-template/helpers"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
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

	Describe("#GetBoolConfig", func() {
		Context("given EXISTING config key", func() {
			It("returns true", func() {
				viper.Set("test.bool_config", true)

				result := helpers.GetBoolConfig("bool_config")

				Expect(result).To(BeTrue())
			})
		})

		Context("given NON-EXISTING config key", func() {
			It("returns false", func() {
				viper.Set("test.bool_config", true)

				result := helpers.GetBoolConfig("invalid")

				Expect(result).To(BeFalse())
			})
		})
	})

	Describe("#GetFloatConfig", func() {
		Context("given EXISTING config key", func() {
			It("returns correct config value", func() {
				viper.Set("test.float_config", 1.0)

				result := helpers.GetFloatConfig("float_config")

				Expect(result).To(Equal(1.0))
			})
		})

		Context("given NON-EXISTING config key", func() {
			It("returns zero", func() {
				viper.Set("test.float_config", 1.0)

				result := helpers.GetFloatConfig("invalid")

				Expect(result).To(BeZero())
			})
		})
	})

	Describe("#GetIntConfig", func() {
		Context("given EXISTING config key", func() {
			It("returns correct config value", func() {
				viper.Set("test.int_config", 1)

				result := helpers.GetIntConfig("int_config")

				Expect(result).To(Equal(1))
			})
		})

		Context("given NON-EXISTING config key", func() {
			It("returns zero", func() {
				viper.Set("test.int_config", 1)

				result := helpers.GetIntConfig("invalid")

				Expect(result).To(BeZero())
			})
		})
	})

	Describe("#GetStringConfig", func() {
		Context("given EXISTING config key", func() {
			It("returns correct config value", func() {
				viper.Set("test.string_config", "some string")

				result := helpers.GetStringConfig("string_config")

				Expect(result).To(Equal("some string"))
			})
		})

		Context("given NON-EXISTING config key", func() {
			It("returns empty config value", func() {
				viper.Set("test.string_config", "some string")

				result := helpers.GetStringConfig("invalid")

				Expect(result).To(BeEmpty())
			})
		})
	})
})
