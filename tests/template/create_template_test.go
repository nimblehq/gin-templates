package template_test

import (
	"os"
	"strings"

	"github.com/nimblehq/gin-templates/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = BeforeEach(func() {
	tests.ChangeDirectory(TemplateGeneratedPath + "/test-gin-templates")
})

var _ = Describe("Create template", func() {
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

	Context("given docker-compose.dev.yml", func() {
		It("contains project name at container_name", func() {
			content := tests.ReadFile("docker-compose.dev.yml")

			isContainProjectName := strings.Contains(content, "container_name: test-gin-templates_db")

			Expect(isContainProjectName).To(BeTrue())
		})

		It("contains project name at postgres db env", func() {
			content := tests.ReadFile("docker-compose.dev.yml")

			isContainProjectName := strings.Contains(content, "- POSTGRES_DB=test-gin-templates_development")

			Expect(isContainProjectName).To(BeTrue())
		})
	})

	Context("given docker-compose.test.yml", func() {
		It("contains project name at container_name", func() {
			content := tests.ReadFile("docker-compose.test.yml")

			isContainProjectName := strings.Contains(content, "container_name: test-gin-templates_db_test")

			Expect(isContainProjectName).To(BeTrue())
		})

		It("contains project name at postgres db env", func() {
			content := tests.ReadFile("docker-compose.test.yml")

			isContainProjectName := strings.Contains(content, "- POSTGRES_DB=test-gin-templates_test")

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

	Context("given bootstrap/database.go", func() {
		It("contains project name at helpers import", func() {
			content := tests.ReadFile("bootstrap/database.go")

			isContainProjectName := strings.Contains(content, `"github.com/nimblehq/test-gin-templates/helpers"`)

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

	Context("given config/app.toml", func() {
		It("contains project name at app_name", func() {
			content := tests.ReadFile("config/app.toml")

			isContainProjectName := strings.Contains(content, `app_name = "test-gin-templates"`)

			Expect(isContainProjectName).To(BeTrue())
		})

		It("contains project name at debug db name", func() {
			content := tests.ReadFile("config/app.toml")

			isContainProjectName := strings.Contains(content, `db_name = "test-gin-templates_development"`)

			Expect(isContainProjectName).To(BeTrue())
		})

		It("contains project name at test db name", func() {
			content := tests.ReadFile("config/app.toml")

			isContainProjectName := strings.Contains(content, `db_name = "test-gin-templates_test"`)

			Expect(isContainProjectName).To(BeTrue())
		})
	})

	Context("given helpers/config_test.go", func() {
		It("contains project name at helpers import", func() {
			content := tests.ReadFile("helpers/config_test.go")

			isContainProjectName := strings.Contains(content, `"github.com/nimblehq/test-gin-templates/helpers"`)

			Expect(isContainProjectName).To(BeTrue())
		})
	})

	Context("given helpers/helpers_suite_test.go", func() {
		It("contains project name at test import", func() {
			content := tests.ReadFile("helpers/helpers_suite_test.go")

			isContainProjectName := strings.Contains(content, `"github.com/nimblehq/test-gin-templates/test"`)

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
