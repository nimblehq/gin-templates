package cmd_test

import (
	"os"

	"github.com/nimblehq/gin-templates/tests"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var templateGeneratedPath string = tests.GetCurrentDirectory()

var _ = BeforeEach(func() {
	tests.ChangeDirectory(templateGeneratedPath)
})

var _ = AfterEach(func() {
	os.RemoveAll(templateGeneratedPath + "/test-gin-templates")
})

var _ = Describe("Create template", func() {
	Context("given the project directory", func() {
		It("contains project name", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(templateGeneratedPath)

			dir, err := os.Getwd()
			if err != nil {
				Fail("Failed to get current directory: " + err.Error())
			}

			expectedContent := "gin-templates/cmd/test-gin-templates"

			Expect(dir).To(ContainSubstring(expectedContent))
		})
	})

	Context("given .env.example", func() {
		It("contains project name at DATABASE_URL", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(templateGeneratedPath)

			content := tests.ReadFile(".env.example")

			expectedContent := "DATABASE_URL=postgres://postgres:postgres@0.0.0.0:5432/test-gin-templates_development?sslmode=disable"

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given .env.test", func() {
		It("contains project name at DATABASE_URL", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(templateGeneratedPath)

			content := tests.ReadFile(".env.test")

			expectedContent := "DATABASE_URL=postgres://postgres:postgres@0.0.0.0:5433/test-gin-templates_test?sslmode=disable"

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given docker-compose.dev.yml", func() {
		It("contains project name at container_name", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(templateGeneratedPath)

			content := tests.ReadFile("docker-compose.dev.yml")

			expectedContent := "container_name: test-gin-templates_db"

			Expect(content).To(ContainSubstring(expectedContent))
		})

		It("contains project name at postgres db env", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(templateGeneratedPath)

			content := tests.ReadFile("docker-compose.dev.yml")

			expectedContent := "- POSTGRES_DB=test-gin-templates_development"

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given docker-compose.test.yml", func() {
		It("contains project name at container_name", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(templateGeneratedPath)

			content := tests.ReadFile("docker-compose.test.yml")

			expectedContent := "container_name: test-gin-templates_db_test"

			Expect(content).To(ContainSubstring(expectedContent))
		})

		It("contains project name at postgres db env", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(templateGeneratedPath)

			content := tests.ReadFile("docker-compose.test.yml")

			expectedContent := "- POSTGRES_DB=test-gin-templates_test"

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given go.mod", func() {
		It("contains project name at module name", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(templateGeneratedPath)

			content := tests.ReadFile("go.mod")

			expectedContent := "module github.com/nimblehq/test-gin-templates"

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given bootstrap/database.go", func() {
		It("contains project name at helpers import", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(templateGeneratedPath)

			content := tests.ReadFile("bootstrap/database.go")

			expectedContent := `"github.com/nimblehq/test-gin-templates/helpers"`

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given bootstrap/router.go", func() {
		It("contains project name at api v1 routers import", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(templateGeneratedPath)

			content := tests.ReadFile("bootstrap/router.go")

			expectedContent := `apiv1router "github.com/nimblehq/test-gin-templates/lib/api/v1/routers"`

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given cmd/api/main.go", func() {
		It("contains project name at bootstrap import", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(templateGeneratedPath)

			content := tests.ReadFile("cmd/api/main.go")

			expectedContent := `"github.com/nimblehq/test-gin-templates/bootstrap"`

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given config/app.toml", func() {
		It("contains project name at app_name", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(templateGeneratedPath)

			content := tests.ReadFile("config/app.toml")

			expectedContent := `app_name = "test-gin-templates"`

			Expect(content).To(ContainSubstring(expectedContent))
		})

		It("contains project name at debug db name", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(templateGeneratedPath)

			content := tests.ReadFile("config/app.toml")

			expectedContent := `db_name = "test-gin-templates_development"`

			Expect(content).To(ContainSubstring(expectedContent))
		})

		It("contains project name at test db name", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(templateGeneratedPath)

			content := tests.ReadFile("config/app.toml")

			expectedContent := `db_name = "test-gin-templates_test"`

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given helpers/config_test.go", func() {
		It("contains project name at helpers import", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(templateGeneratedPath)

			content := tests.ReadFile("helpers/config_test.go")

			expectedContent := `"github.com/nimblehq/test-gin-templates/helpers"`

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given helpers/helpers_suite_test.go", func() {
		It("contains project name at test import", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(templateGeneratedPath)

			content := tests.ReadFile("helpers/helpers_suite_test.go")

			expectedContent := `"github.com/nimblehq/test-gin-templates/test"`

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given lib/api/v1/controllers/controllers_suite_test.go", func() {
		It("contains project name at test import", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(templateGeneratedPath)

			content := tests.ReadFile("lib/api/v1/controllers/controllers_suite_test.go")

			expectedContent := `"github.com/nimblehq/test-gin-templates/test"`

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given lib/api/v1/controllers/health_test.go", func() {
		It("contains project name at test import", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(templateGeneratedPath)

			content := tests.ReadFile("lib/api/v1/controllers/health_test.go")

			expectedContent := `"github.com/nimblehq/test-gin-templates/test"`

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given lib/api/v1/routers/router.go", func() {
		It("contains project name at api v1 controllers import", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(templateGeneratedPath)

			content := tests.ReadFile("lib/api/v1/routers/router.go")

			expectedContent := `"github.com/nimblehq/test-gin-templates/lib/api/v1/controllers"`

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given test/test.go", func() {
		It("contains project name at bootstrap import", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(templateGeneratedPath)

			content := tests.ReadFile("test/test.go")

			expectedContent := `"github.com/nimblehq/test-gin-templates/bootstrap"`

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})
})
