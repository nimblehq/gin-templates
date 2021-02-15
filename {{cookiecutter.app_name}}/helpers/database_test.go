package helpers_test

import (
	"errors"

	"github.com/nimblehq/{{cookiecutter.app_name}}/helpers"
	. "github.com/onsi/ginkgo"

	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var _ = Describe("Database", func() {
	Describe("#GetDatabaseConnection", func() {
		Context("given *gorm.DB type to viper config", func() {
			It("returns database connection", func() {
				viper.Set("database", &gorm.DB{})

				result, err := helpers.GetDBConnection()

				Expect(result).To(BeAssignableToTypeOf(&gorm.DB{}))
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("given INVALID type to viper config", func() {
			It("returns an error", func() {
				viper.Set("database", "invalid")

				result, err := helpers.GetDBConnection()

				Expect(result).To(BeNil())
				Expect(err).To(MatchError(errors.New("Failed to get database connection: connection is not type of *gorm.DB")))
			})
		})

		Context("given NOTHING to viper config", func() {
			It("returns an error", func() {
				result, err := helpers.GetDBConnection()

				Expect(result).To(BeNil())
				Expect(err).To(MatchError(errors.New("Failed to get database connection: connection is not type of *gorm.DB")))
			})
		})
	})
})
