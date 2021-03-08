package helpers_test

import (
	"testing"

	"github.com/nimblehq/{{cookiecutter.app_name}}/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHelpers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Helpers Suite")
}

var _ = BeforeSuite(func() {
	test.SetupTestEnvironment()
})
