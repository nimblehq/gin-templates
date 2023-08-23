package cmd_test

import (
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
		It("contains project name at database import", func() {
			cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)

			content := tests.ReadFile("bootstrap/database.go")

			expectedContent := `"github.com/nimblehq/test-gin-templates/database"`

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

			expectedContent := "github.com/sirupsen/logrus"

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

		Context("given database/database.go", func() {
			It("contains project name at helpers import", func() {
				cookiecutter := tests.Cookiecutter{AppName: "test-gin-templates"}
				cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)

				content := tests.ReadFile("database/database.go")

				expectedContent := `"github.com/nimblehq/test-gin-templates/helpers"`

				Expect(content).To(ContainSubstring(expectedContent))
			})
		})

		It("contains logrus package import in database/database.go", func() {
			cookiecutter := tests.Cookiecutter{
				AppName:   "test-gin-templates",
				UseLogrus: tests.Yes,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("database/database.go")

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
			files, err := os.ReadDir("deploy/heroku")
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

		It("contains heroku instruction in README", func() {
			cookiecutter := tests.Cookiecutter{
				AppName:   "test-gin-templates",
				UseHeroku: tests.Yes,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("README.md")

			expectedContent := "Deploy to Heroku with Terraform"

			Expect(content).To(ContainSubstring(expectedContent))
		})

		It("contains heroku files in .gitignore", func() {
			cookiecutter := tests.Cookiecutter{
				AppName:   "test-gin-templates",
				UseHeroku: tests.Yes,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile(".gitignore")

			expectedContent := "deploy/**/.terraform/"

			Expect(content).To(ContainSubstring(expectedContent))
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

		It("does NOT contains heroku instruction in README", func() {
			cookiecutter := tests.Cookiecutter{
				AppName:   "test-gin-templates",
				UseHeroku: tests.No,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("README.md")

			expectedContent := "Deploy to Heroku with Terraform"

			Expect(content).NotTo(ContainSubstring(expectedContent))
		})

		It("does NOT contains heroku files in .gitignore", func() {
			cookiecutter := tests.Cookiecutter{
				AppName:   "test-gin-templates",
				UseHeroku: tests.No,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile(".gitignore")

			expectedContent := "deploy/**/.terraform/"

			Expect(content).NotTo(ContainSubstring(expectedContent))
		})
	})

	Context("given Web variant", func() {
		It("contains lib/web folder", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
				Variant: tests.Web,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			_, err := os.Stat("lib/web")

			Expect(os.IsNotExist(err)).To(BeFalse())
		})

		It("contains .eslintrc.json file", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
				Variant: tests.Web,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			_, err := os.Stat(".eslintrc.json")

			Expect(os.IsNotExist(err)).To(BeFalse())
		})

		It("contains .npmrc file", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
				Variant: tests.Web,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			_, err := os.Stat(".npmrc")

			Expect(os.IsNotExist(err)).To(BeFalse())
		})

		It("contains package.json file", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
				Variant: tests.Web,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			_, err := os.Stat("package.json")

			Expect(os.IsNotExist(err)).To(BeFalse())
		})

		It("contains snowpack.config.js file", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
				Variant: tests.Web,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			_, err := os.Stat("snowpack.config.js")

			Expect(os.IsNotExist(err)).To(BeFalse())
		})

		It("contains postcss.config.js file", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
				Variant: tests.Web,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			_, err := os.Stat("postcss.config.js")

			Expect(os.IsNotExist(err)).To(BeFalse())
		})

		It("contains tsconfig.json file", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
				Variant: tests.Web,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			_, err := os.Stat("tsconfig.json")

			Expect(os.IsNotExist(err)).To(BeFalse())
		})

		It("contains snowpack: npm run dev in Procfile.dev", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
				Variant: tests.Web,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("Procfile.dev")

			expectedContent := "snowpack: npm run dev"

			Expect(content).To(ContainSubstring(expectedContent))
		})

		It("contains the webrouter import in bootstrap/router.go", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
				Variant: tests.Web,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("bootstrap/router.go")

			expectedContent := "webrouter \"github.com/nimblehq/test-gin-templates/lib/web/routers\""

			Expect(content).To(ContainSubstring(expectedContent))
		})

		It("contains npm install in Makefile", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
				Variant: tests.Web,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("Makefile")

			expectedContent := "npm install"

			Expect(content).To(ContainSubstring(expectedContent))
		})

		It("contains Node 14 in README.md", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
				Variant: tests.Web,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("README.md")

			expectedContent := "[Node - 14](https://nodejs.org/en/)"

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given NO Web variant", func() {
		It("does NOT contain lib/web folder", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
				Variant: tests.API,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			_, err := os.Stat("lib/web")

			Expect(os.IsNotExist(err)).To(BeTrue())
		})

		It("does NOT contain .npmrc file", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
				Variant: tests.API,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			_, err := os.Stat(".npmrc")

			Expect(os.IsNotExist(err)).To(BeTrue())
		})

		It("does NOT contain snowpack.config.js file", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
				Variant: tests.API,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			_, err := os.Stat("snowpack.config.js")

			Expect(os.IsNotExist(err)).To(BeTrue())
		})

		It("does NOT contain postcss.config.js file", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
				Variant: tests.API,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			_, err := os.Stat("postcss.config.js")

			Expect(os.IsNotExist(err)).To(BeTrue())
		})

		It("does NOT contain tsconfig.json file", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
				Variant: tests.API,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			_, err := os.Stat("tsconfig.json")

			Expect(os.IsNotExist(err)).To(BeTrue())
		})

		It("does NOT contain snowpack: npm run dev in Procfile.dev", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
				Variant: tests.API,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("Procfile.dev")

			expectedContent := "snowpack: npm run dev"

			Expect(content).NotTo(ContainSubstring(expectedContent))
		})

		It("does NOT contain the webrouter import in bootstrap/router.go", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
				Variant: tests.API,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("bootstrap/router.go")

			expectedContent := "webrouter \"github.com/nimblehq/test-gin-templates/lib/web/routers\""

			Expect(content).NotTo(ContainSubstring(expectedContent))
		})

		It("does NOT contain Node 14 in README.md", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
				Variant: tests.API,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("README.md")

			expectedContent := "[Node - 14](https://nodejs.org/en/)"

			Expect(content).NotTo(ContainSubstring(expectedContent))
		})
	})

	Context("given bootstrap add-on", func() {
		It("contains lib/web/assets/stylesheets/vendors/bootstrap folder", func() {
			cookiecutter := tests.Cookiecutter{
				Variant:  tests.Web,
				CssAddon: tests.Bootstrap,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			_, err := os.Stat("lib/web/assets/stylesheets/vendors/bootstrap")

			Expect(os.IsNotExist(err)).To(BeFalse())
		})

		It("contains bootstrap package in package.json", func() {
			cookiecutter := tests.Cookiecutter{
				Variant:  tests.Web,
				CssAddon: tests.Bootstrap,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("package.json")

			expectedContent := "\"bootstrap\": \"5.0.2\""

			Expect(content).To(ContainSubstring(expectedContent))
		})

		It("contains bootstrap scss import in lib/web/assets/stylesheets/application.scss", func() {
			cookiecutter := tests.Cookiecutter{
				Variant:  tests.Web,
				CssAddon: tests.Bootstrap,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("lib/web/assets/stylesheets/application.scss")

			expectedContent := "@import \"./vendors/bootstrap\";"

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given NO bootstrap add-on", func() {
		It("does NOT contain lib/web/assets/stylesheets/vendors/bootstrap folder", func() {
			cookiecutter := tests.Cookiecutter{
				Variant:  tests.Web,
				CssAddon: tests.None,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			_, err := os.Stat("lib/web/assets/stylesheets/vendors/bootstrap")

			Expect(os.IsNotExist(err)).To(BeTrue())
		})

		It("does NOT contain bootstrap package in package.json", func() {
			cookiecutter := tests.Cookiecutter{
				Variant:  tests.Web,
				CssAddon: tests.None,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("package.json")

			expectedContent := "\"bootstrap\":"

			Expect(content).NotTo(ContainSubstring(expectedContent))
		})

		It("does NOT contain bootstrap scss import in lib/web/assets/stylesheets/application.scss", func() {
			cookiecutter := tests.Cookiecutter{
				Variant:  tests.Web,
				CssAddon: tests.None,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("lib/web/assets/stylesheets/application.scss")

			expectedContent := "@import \"./vendors/bootstrap\";"

			Expect(content).NotTo(ContainSubstring(expectedContent))
		})
	})

	Context("given tailwind add-on", func() {
		It("contains lib/web/assets/stylesheets/vendors/tailwind folder", func() {
			cookiecutter := tests.Cookiecutter{
				Variant:  tests.Web,
				CssAddon: tests.Tailwind,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			_, err := os.Stat("lib/web/assets/stylesheets/vendors/tailwind")

			Expect(os.IsNotExist(err)).To(BeFalse())
		})

		It("contains tailwind.config.js file", func() {
			cookiecutter := tests.Cookiecutter{
				Variant:  tests.Web,
				CssAddon: tests.Tailwind,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			_, err := os.Stat("tailwind.config.js")

			Expect(os.IsNotExist(err)).To(BeFalse())
		})

		It("contains tailwind package in package.json", func() {
			cookiecutter := tests.Cookiecutter{
				Variant:  tests.Web,
				CssAddon: tests.Tailwind,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("package.json")

			expectedContent := "\"tailwindcss\": \"2.2.7\""

			Expect(content).To(ContainSubstring(expectedContent))
		})

		It("contains tailwind scss import in lib/web/assets/stylesheets/application.scss", func() {
			cookiecutter := tests.Cookiecutter{
				Variant:  tests.Web,
				CssAddon: tests.Tailwind,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("lib/web/assets/stylesheets/application.scss")

			expectedContent := "@import \"./vendors/tailwind\";"

			Expect(content).To(ContainSubstring(expectedContent))
		})

		It("contains tailwind requirement in postcss.config.js", func() {
			cookiecutter := tests.Cookiecutter{
				Variant:  tests.Web,
				CssAddon: tests.Tailwind,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("postcss.config.js")

			expectedContent := "require(\"tailwindcss\")"

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})

	Context("given NO tailwind add-on", func() {
		It("does NOT contain lib/web/assets/stylesheets/vendors/tailwind folder", func() {
			cookiecutter := tests.Cookiecutter{
				Variant:  tests.Web,
				CssAddon: tests.None,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			_, err := os.Stat("lib/web/assets/stylesheets/vendors/tailwind")

			Expect(os.IsNotExist(err)).To(BeTrue())
		})

		It("does NOT contain tailwind.config.js file", func() {
			cookiecutter := tests.Cookiecutter{
				Variant:  tests.Web,
				CssAddon: tests.None,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			_, err := os.Stat("tailwind.config.js")

			Expect(os.IsNotExist(err)).To(BeTrue())
		})

		It("does NOT contain tailwind package in package.json", func() {
			cookiecutter := tests.Cookiecutter{
				Variant:  tests.Web,
				CssAddon: tests.None,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("package.json")

			expectedContent := "\"tailwindcss\":"

			Expect(content).NotTo(ContainSubstring(expectedContent))
		})

		It("does NOT contain tailwind scss import in lib/web/assets/stylesheets/application.scss", func() {
			cookiecutter := tests.Cookiecutter{
				Variant:  tests.Web,
				CssAddon: tests.None,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("lib/web/assets/stylesheets/application.scss")

			expectedContent := "@import \"./vendors/tailwind\";"

			Expect(content).NotTo(ContainSubstring(expectedContent))
		})

		It("does NOT contain tailwind requirement in postcss.config.js", func() {
			cookiecutter := tests.Cookiecutter{
				Variant:  tests.Web,
				CssAddon: tests.None,
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("postcss.config.js")

			expectedContent := "require(\"tailwindcss\")"

			Expect(content).NotTo(ContainSubstring(expectedContent))
		})
	})

	Context("given openapi add-on", func() {
		Context("given only Web variant", func() {
			It("does NOT contains openapi requirement in .eslintrc.json", func() {
				cookiecutter := tests.Cookiecutter{
					AppName: "test-gin-templates",
					Variant: tests.Web,
				}
				cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
				content := tests.ReadFile(".eslintrc.json")

				expectedContent := `plugin:yml/recommended`

				Expect(content).NotTo(ContainSubstring(expectedContent))
			})
		})

		Context("given only API variant", func() {
			It("contains openapi requirement in .eslintrc.json", func() {
				cookiecutter := tests.Cookiecutter{
					AppName: "test-gin-templates",
					Variant: tests.API,
				}
				cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
				content := tests.ReadFile(".eslintrc.json")

				expectedContent := `plugin:yml/recommended`

				Expect(content).To(ContainSubstring(expectedContent))
			})
		})

		Context("given both variant", func() {
			It("contains openapi requirement in .eslintrc.json", func() {
				cookiecutter := tests.Cookiecutter{
					AppName: "test-gin-templates",
					Variant: tests.Both,
				}
				cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
				content := tests.ReadFile(".eslintrc.json")

				expectedContent := `plugin:yml/recommended`

				Expect(content).To(ContainSubstring(expectedContent))
			})
		})

		It("contains docs/openapi folder", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			_, err := os.Stat("docs/openapi")

			Expect(os.IsNotExist(err)).To(BeFalse())
		})

		It("contains openapi instruction in README", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("README.md")

			expectedContent := "Generate API documentation"

			Expect(content).To(ContainSubstring(expectedContent))
		})

		It("contains .dockerignore file", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			_, err := os.Stat(".dockerignore")

			Expect(os.IsNotExist(err)).To(BeFalse())
		})

		It("contains .eslintignore file", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			_, err := os.Stat(".eslintignore")

			Expect(os.IsNotExist(err)).To(BeFalse())
		})

		It("contains .github/workflows/lint_docs.yml file", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			_, err := os.Stat(".github/workflows/lint_docs.yml")

			Expect(os.IsNotExist(err)).To(BeFalse())
		})

		It("contains openapi requirement in .gitignore", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile(".gitignore")

			expectedContent := "/public/openapi.yml"

			Expect(content).To(ContainSubstring(expectedContent))
		})

		It("contains .spectral.yaml file", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			_, err := os.Stat(".spectral.yaml")

			Expect(os.IsNotExist(err)).To(BeFalse())
		})

		It("contains openapi requirement in .tool-versions", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile(".tool-versions")

			expectedContent := "nodejs 18.15.0"

			Expect(content).To(ContainSubstring(expectedContent))
		})

		It("contains openapi requirement in Makefile", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("Makefile")

			expectedContent := "npm run build:docs"

			Expect(content).To(ContainSubstring(expectedContent))
		})

		It("contains openapi requirement in package.json", func() {
			cookiecutter := tests.Cookiecutter{
				AppName: "test-gin-templates",
			}
			cookiecutter.CreateProjectFromGinTemplate(currentTemplatePath)
			content := tests.ReadFile("package.json")

			expectedContent := `"build:docs": "swagger-cli bundle docs/openapi/openapi.yml --outfile public/openapi.yml --type yaml"`

			Expect(content).To(ContainSubstring(expectedContent))
		})
	})
})
