package cmd_test

import (
	"io/ioutil"
	"os"

	"github.com/nimblehq/gin-templates/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var currentTemplatePath string = tests.GetCurrentDirectory()

var _ = BeforeEach(func() {
	tests.ChangeDirectory(currentTemplatePath)
})

var _ = AfterEach(func() {
	os.RemoveAll(currentTemplatePath + "/test-gin-templates")
})

var _ = Describe("Create template", func() {
	Context("given the project directory", func() {
		It("contains project name", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)

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
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)

			content := tests.ReadFile(".env.example")

			expectedContent := "DATABASE_URL=postgres://postgres:postgres@0.0.0.0:5432/test-gin-templates_development?sslmode=disable"

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given .env.test", func() {
		It("contains project name at DATABASE_URL", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)

			content := tests.ReadFile(".env.test")

			expectedContent := "DATABASE_URL=postgres://postgres:postgres@0.0.0.0:5433/test-gin-templates_test?sslmode=disable"

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given docker-compose.dev.yml", func() {
		It("contains project name at container_name", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)

			content := tests.ReadFile("docker-compose.dev.yml")

			expectedContent := "container_name: test-gin-templates_db"

			Expect(content).To(ContainSubstring(expectedContent))
		})

		It("contains project name at postgres db env", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)

			content := tests.ReadFile("docker-compose.dev.yml")

			expectedContent := "- POSTGRES_DB=test-gin-templates_development"

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given docker-compose.test.yml", func() {
		It("contains project name at container_name", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)

			content := tests.ReadFile("docker-compose.test.yml")

			expectedContent := "container_name: test-gin-templates_db_test"

			Expect(content).To(ContainSubstring(expectedContent))
		})

		It("contains project name at postgres db env", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)

			content := tests.ReadFile("docker-compose.test.yml")

			expectedContent := "- POSTGRES_DB=test-gin-templates_test"

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given go.mod", func() {
		It("contains project name at module name", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)

			content := tests.ReadFile("go.mod")

			expectedContent := "module github.com/nimblehq/test-gin-templates"

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given bootstrap/database.go", func() {
		It("contains project name at helpers import", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)

			content := tests.ReadFile("bootstrap/database.go")

			expectedContent := `"github.com/nimblehq/test-gin-templates/helpers"`

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given bootstrap/router.go", func() {
		It("contains project name at api v1 routers import", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)

			content := tests.ReadFile("bootstrap/router.go")

			expectedContent := `apiv1router "github.com/nimblehq/test-gin-templates/lib/api/v1/routers"`

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given cmd/api/main.go", func() {
		It("contains project name at bootstrap import", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)

			content := tests.ReadFile("cmd/api/main.go")

			expectedContent := `"github.com/nimblehq/test-gin-templates/bootstrap"`

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given config/app.toml", func() {
		It("contains project name at app_name", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)

			content := tests.ReadFile("config/app.toml")

			expectedContent := `app_name = "test-gin-templates"`

			Expect(content).To(ContainSubstring(expectedContent))
		})

		It("contains project name at debug db name", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)

			content := tests.ReadFile("config/app.toml")

			expectedContent := `db_name = "test-gin-templates_development"`

			Expect(content).To(ContainSubstring(expectedContent))
		})

		It("contains project name at test db name", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)

			content := tests.ReadFile("config/app.toml")

			expectedContent := `db_name = "test-gin-templates_test"`

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given helpers/config_test.go", func() {
		It("contains project name at helpers import", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)

			content := tests.ReadFile("helpers/config_test.go")

			expectedContent := `"github.com/nimblehq/test-gin-templates/helpers"`

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given helpers/helpers_suite_test.go", func() {
		It("contains project name at test import", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)

			content := tests.ReadFile("helpers/helpers_suite_test.go")

			expectedContent := `"github.com/nimblehq/test-gin-templates/test"`

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given lib/api/v1/controllers/controllers_suite_test.go", func() {
		It("contains project name at test import", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)

			content := tests.ReadFile("lib/api/v1/controllers/controllers_suite_test.go")

			expectedContent := `"github.com/nimblehq/test-gin-templates/test"`

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given lib/api/v1/controllers/health_test.go", func() {
		It("contains project name at test import", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)

			content := tests.ReadFile("lib/api/v1/controllers/health_test.go")

			expectedContent := `"github.com/nimblehq/test-gin-templates/test"`

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given lib/api/v1/routers/router.go", func() {
		It("contains project name at api v1 controllers import", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)

			content := tests.ReadFile("lib/api/v1/routers/router.go")

			expectedContent := `"github.com/nimblehq/test-gin-templates/lib/api/v1/controllers"`

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given test/test.go", func() {
		It("contains project name at bootstrap import", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)

			content := tests.ReadFile("test/test.go")

			expectedContent := `"github.com/nimblehq/test-gin-templates/bootstrap"`

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given logrus add-on", func() {
		It("contains helpers/log folder", func() {
			cookiecutter := tests.Cookiecutter{
				AppName:   "test-gin-templates",
				UseLogrus: tests.Yes,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			_, err := os.Stat("helpers/log")

			Expect(os.IsNotExist(err)).To(BeFalse())
		})

		It("contains logrus package in go.mod", func() {
			cookiecutter := tests.Cookiecutter{
				AppName:   "test-gin-templates",
				UseLogrus: tests.Yes,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("go.mod")

			expectedContent := "github.com/sirupsen/logrus v1.8.1"

			Expect(content).To(ContainSubstring(expectedContent))
		})

		It("contains logrus package import in bootstrap/config.go", func() {
			cookiecutter := tests.Cookiecutter{
				AppName:   "test-gin-templates",
				UseLogrus: tests.Yes,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("bootstrap/config.go")

			expectedContent := "github.com/nimblehq/test-gin-templates/helpers/log"

			Expect(content).To(ContainSubstring(expectedContent))
		})

		It("contains logrus package import in bootstrap/database.go", func() {
			cookiecutter := tests.Cookiecutter{
				AppName:   "test-gin-templates",
				UseLogrus: tests.Yes,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("bootstrap/database.go")

			expectedContent := "github.com/nimblehq/test-gin-templates/helpers/log"

			Expect(content).To(ContainSubstring(expectedContent))
		})

		It("contains logrus package import in cmd/api/main.go", func() {
			cookiecutter := tests.Cookiecutter{
				AppName:   "test-gin-templates",
				UseLogrus: tests.Yes,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("cmd/api/main.go")

			expectedContent := "github.com/nimblehq/test-gin-templates/helpers/log"

			Expect(content).To(ContainSubstring(expectedContent))
		})

		It("contains logrus package import in test/test.go", func() {
			cookiecutter := tests.Cookiecutter{
				AppName:   "test-gin-templates",
				UseLogrus: tests.Yes,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("test/test.go")

			expectedContent := "github.com/nimblehq/test-gin-templates/helpers/log"

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given NO logrus add-on", func() {
		It("does NOT contain helpers/log folder", func() {
			cookiecutter := tests.Cookiecutter{
				AppName:   "test-gin-templates",
				UseLogrus: tests.No,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			_, err := os.Stat("helpers/log")

			Expect(os.IsNotExist(err)).To(BeTrue())
		})

		It("does NOT contain logrus package in go.mod", func() {
			cookiecutter := tests.Cookiecutter{
				AppName:   "test-gin-templates",
				UseLogrus: tests.No,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("go.mod")

			expectedContent := "github.com/sirupsen/logrus v1.8.1"

			Expect(content).NotTo(ContainSubstring(expectedContent))
		})

		It("does NOT contain logrus package import in bootstrap/config.go", func() {
			cookiecutter := tests.Cookiecutter{
				AppName:   "test-gin-templates",
				UseLogrus: tests.No,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("bootstrap/config.go")

			expectedContent := "github.com/nimblehq/test-gin-templates/helpers/log"

			Expect(content).NotTo(ContainSubstring(expectedContent))
		})

		It("does NOT contain logrus package import in bootstrap/database.go", func() {
			cookiecutter := tests.Cookiecutter{
				AppName:   "test-gin-templates",
				UseLogrus: tests.No,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("bootstrap/database.go")

			expectedContent := "github.com/nimblehq/test-gin-templates/helpers/log"

			Expect(content).NotTo(ContainSubstring(expectedContent))
		})

		It("does NOT contain logrus package import in cmd/api/main.go", func() {
			cookiecutter := tests.Cookiecutter{
				AppName:   "test-gin-templates",
				UseLogrus: tests.No,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("cmd/api/main.go")

			expectedContent := "github.com/nimblehq/test-gin-templates/helpers/log"

			Expect(content).NotTo(ContainSubstring(expectedContent))
		})

		It("does NOT contain logrus package import in test/test.go", func() {
			cookiecutter := tests.Cookiecutter{
				AppName:   "test-gin-templates",
				UseLogrus: tests.No,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("test/test.go")

			expectedContent := "github.com/nimblehq/test-gin-templates/helpers/log"

			Expect(content).NotTo(ContainSubstring(expectedContent))
		})
	})

	Context("given heroku add-on", func() {
		It("contains deploy/heroku folder", func() {
			cookiecutter := tests.Cookiecutter{
				AppName:   "test-gin-templates",
				UseHeroku: tests.Yes,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			_, err := os.Stat("deploy/heroku")

			Expect(os.IsNotExist(err)).To(BeFalse())
		})

		It("contains valid files in deploy/heroku folder", func() {
			cookiecutter := tests.Cookiecutter{
				AppName:   "test-gin-templates",
				UseHeroku: tests.Yes,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			files, err := ioutil.ReadDir("deploy/heroku")
			if err != nil {
				Fail("Failed to read directory: " + err.Error())
			}

			expectedFiles := []string{
				"config.tf",
				"main.tf",
				"terraform.tfvars.sample",
				"variables.tf",
			}

			Expect(len(files)).To(Equal(len(expectedFiles)))

			for k, f := range files {
				Expect(f.Name()).To(Equal(expectedFiles[k]))
			}
		})
	})

	Context("given NO heroku add-on", func() {
		It("does NOT contains deploy/heroku folder", func() {
			cookiecutter := tests.Cookiecutter{
				AppName:   "test-gin-templates",
				UseHeroku: tests.No,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			_, err := os.Stat("deploy/heroku")

			Expect(os.IsNotExist(err)).To(BeTrue())
		})
	})
})
