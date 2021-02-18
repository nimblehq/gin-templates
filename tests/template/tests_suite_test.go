package template_test

import (
	"os"
	"testing"

	"github.com/nimblehq/gin-templates/tests"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var TemplateGeneratedPath string

func TestTemplate(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Template Suite")
}

var _ = BeforeSuite(func() {
	TemplateGeneratedPath = tests.CreateTemplate()
})

var _ = AfterSuite(func() {
	os.RemoveAll(TemplateGeneratedPath + "/test-gin-templates")
})
