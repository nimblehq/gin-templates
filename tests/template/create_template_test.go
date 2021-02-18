package template_test

import (
	"os"
	"strings"

	"github.com/nimblehq/gin-templates/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = BeforeEach(func() {
	tests.ChangeDirectory(ProjectGeneratedPath + "/test-gin-templates")
})

var _ = Describe("Create", func() {
	Context("given current directory", func() {
		It("contains project name", func() {
			dir, err := os.Getwd()
			if err != nil {
				Fail("Failed to get current directory: " + err.Error())
			}

			isContainProjectName := strings.Contains(dir, "gin-templates/tests/template/test-gin-templates")

			Expect(isContainProjectName).To(BeTrue())
		})
	})

	Context("given go.mod", func() {
		It("contains project name at module name", func() {
			content := tests.ReadFile("go.mod")

			isContainProjectName := strings.Contains(content, "module github.com/nimblehq/test-gin-templates")

			Expect(isContainProjectName).To(BeTrue())
		})
	})

	Context("given bootstrap/router.go", func() {
		It("contains project name at api v1 routers import", func() {
			content := tests.ReadFile("bootstrap/router.go")

			isContainProjectName := strings.Contains(content, `apiv1router "github.com/nimblehq/test-gin-templates/lib/api/v1/routers"`)

			Expect(isContainProjectName).To(BeTrue())
		})
	})

	Context("given cmd/api/main.go", func() {
		It("contains project name at bootstrap import", func() {
			content := tests.ReadFile("cmd/api/main.go")

			isContainProjectName := strings.Contains(content, `"github.com/nimblehq/test-gin-templates/bootstrap"`)

			Expect(isContainProjectName).To(BeTrue())
		})
	})

	Context("given lib/api/v1/controllers/controllers_suite_test.go", func() {
		It("contains project name at test import", func() {
			content := tests.ReadFile("lib/api/v1/controllers/controllers_suite_test.go")

			isContainProjectName := strings.Contains(content, `"github.com/nimblehq/test-gin-templates/test"`)

			Expect(isContainProjectName).To(BeTrue())
		})
	})

	Context("given lib/api/v1/controllers/health_test.go", func() {
		It("contains project name at test import", func() {
			content := tests.ReadFile("lib/api/v1/controllers/health_test.go")

			isContainProjectName := strings.Contains(content, `"github.com/nimblehq/test-gin-templates/test"`)

			Expect(isContainProjectName).To(BeTrue())
		})
	})

	Context("given lib/api/v1/routers/router.go", func() {
		It("contains project name at api v1 controllers import", func() {
			content := tests.ReadFile("lib/api/v1/routers/router.go")

			isContainProjectName := strings.Contains(content, `"github.com/nimblehq/test-gin-templates/lib/api/v1/controllers"`)

			Expect(isContainProjectName).To(BeTrue())
		})
	})

	Context("given test/test.go", func() {
		It("contains project name at bootstrap import", func() {
			content := tests.ReadFile("test/test.go")

			isContainProjectName := strings.Contains(content, `"github.com/nimblehq/test-gin-templates/bootstrap"`)

			Expect(isContainProjectName).To(BeTrue())
		})
	})
})
