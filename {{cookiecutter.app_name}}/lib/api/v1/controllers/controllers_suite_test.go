package controllers_test

import (
	"testing"

	"github.com/nimblehq/{{cookiecutter.app_name}}/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllers Suite")
}

var _ = BeforeSuite(func() {
	test.SetupTestEnvironment()
})
