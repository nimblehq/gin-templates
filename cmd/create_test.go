package cmd_test

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
	Context("given the project directory", func() {
		It("contains project name", func() {
			dir, err := os.Getwd()
			if err != nil {
				Fail("Failed to get current directory: " + err.Error())
			}

			directoryContainProjectName := strings.Contains(dir, "gin-templates/cmd/test-gin-templates")

			Expect(directoryContainProjectName).To(BeTrue())
		})
	})

	Context("given docker-compose.dev.yml", func() {
		It("contains project name at container_name", func() {
			content := tests.ReadFile("docker-compose.dev.yml")

			fileContainProjectName := strings.Contains(content, "container_name: test-gin-templates_db")

			Expect(fileContainProjectName).To(BeTrue())
		})

		It("contains project name at postgres db env", func() {
			content := tests.ReadFile("docker-compose.dev.yml")

			fileContainProjectName := strings.Contains(content, "- POSTGRES_DB=test-gin-templates_development")

			Expect(fileContainProjectName).To(BeTrue())
		})
	})

	Context("given docker-compose.test.yml", func() {
		It("contains project name at container_name", func() {
			content := tests.ReadFile("docker-compose.test.yml")

			fileContainProjectName := strings.Contains(content, "container_name: test-gin-templates_db_test")

			Expect(fileContainProjectName).To(BeTrue())
		})

		It("contains project name at postgres db env", func() {
			content := tests.ReadFile("docker-compose.test.yml")

			fileContainProjectName := strings.Contains(content, "- POSTGRES_DB=test-gin-templates_test")

			Expect(fileContainProjectName).To(BeTrue())
		})
	})

	Context("given go.mod", func() {
		It("contains project name at module name", func() {
			content := tests.ReadFile("go.mod")

			fileContainProjectName := strings.Contains(content, "module github.com/nimblehq/test-gin-templates")

			Expect(fileContainProjectName).To(BeTrue())
		})
	})

	Context("given bootstrap/database.go", func() {
		It("contains project name at helpers import", func() {
			content := tests.ReadFile("bootstrap/database.go")

			fileContainProjectName := strings.Contains(content, `"github.com/nimblehq/test-gin-templates/helpers"`)

			Expect(fileContainProjectName).To(BeTrue())
		})
	})

	Context("given bootstrap/router.go", func() {
		It("contains project name at api v1 routers import", func() {
			content := tests.ReadFile("bootstrap/router.go")

			fileContainProjectName := strings.Contains(content, `apiv1router "github.com/nimblehq/test-gin-templates/lib/api/v1/routers"`)

			Expect(fileContainProjectName).To(BeTrue())
		})
	})

	Context("given cmd/api/main.go", func() {
		It("contains project name at bootstrap import", func() {
			content := tests.ReadFile("cmd/api/main.go")

			fileContainProjectName := strings.Contains(content, `"github.com/nimblehq/test-gin-templates/bootstrap"`)

			Expect(fileContainProjectName).To(BeTrue())
		})
	})

	Context("given config/app.toml", func() {
		It("contains project name at app_name", func() {
			content := tests.ReadFile("config/app.toml")

			fileContainProjectName := strings.Contains(content, `app_name = "test-gin-templates"`)

			Expect(fileContainProjectName).To(BeTrue())
		})

		It("contains project name at debug db name", func() {
			content := tests.ReadFile("config/app.toml")

			fileContainProjectName := strings.Contains(content, `db_name = "test-gin-templates_development"`)

			Expect(fileContainProjectName).To(BeTrue())
		})

		It("contains project name at test db name", func() {
			content := tests.ReadFile("config/app.toml")

			fileContainProjectName := strings.Contains(content, `db_name = "test-gin-templates_test"`)

			Expect(fileContainProjectName).To(BeTrue())
		})
	})

	Context("given helpers/config_test.go", func() {
		It("contains project name at helpers import", func() {
			content := tests.ReadFile("helpers/config_test.go")

			fileContainProjectName := strings.Contains(content, `"github.com/nimblehq/test-gin-templates/helpers"`)

			Expect(fileContainProjectName).To(BeTrue())
		})
	})

	Context("given helpers/helpers_suite_test.go", func() {
		It("contains project name at test import", func() {
			content := tests.ReadFile("helpers/helpers_suite_test.go")

			fileContainProjectName := strings.Contains(content, `"github.com/nimblehq/test-gin-templates/test"`)

			Expect(fileContainProjectName).To(BeTrue())
		})
	})

	Context("given lib/api/v1/controllers/controllers_suite_test.go", func() {
		It("contains project name at test import", func() {
			content := tests.ReadFile("lib/api/v1/controllers/controllers_suite_test.go")

			fileContainProjectName := strings.Contains(content, `"github.com/nimblehq/test-gin-templates/test"`)

			Expect(fileContainProjectName).To(BeTrue())
		})
	})

	Context("given lib/api/v1/controllers/health_test.go", func() {
		It("contains project name at test import", func() {
			content := tests.ReadFile("lib/api/v1/controllers/health_test.go")

			fileContainProjectName := strings.Contains(content, `"github.com/nimblehq/test-gin-templates/test"`)

			Expect(fileContainProjectName).To(BeTrue())
		})
	})

	Context("given lib/api/v1/routers/router.go", func() {
		It("contains project name at api v1 controllers import", func() {
			content := tests.ReadFile("lib/api/v1/routers/router.go")

			fileContainProjectName := strings.Contains(content, `"github.com/nimblehq/test-gin-templates/lib/api/v1/controllers"`)

			Expect(fileContainProjectName).To(BeTrue())
		})
	})

	Context("given test/test.go", func() {
		It("contains project name at bootstrap import", func() {
			content := tests.ReadFile("test/test.go")

			fileContainProjectName := strings.Contains(content, `"github.com/nimblehq/test-gin-templates/bootstrap"`)

			Expect(fileContainProjectName).To(BeTrue())
		})
	})
})
