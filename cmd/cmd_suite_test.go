package cmd_test

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
	tests.DownloadGinTemplate()
	tests.BuildGinTemplate()
	tests.RemoveCookiecuttersCache()
	TemplateGeneratedPath = tests.CreateProjectFromGinTemplate()
})

var _ = AfterSuite(func() {
	os.RemoveAll(TemplateGeneratedPath + "/test-gin-templates")
})
