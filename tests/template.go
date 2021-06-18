package tests

import (
	"io"
	"log"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
)

func DownloadGinTemplate() {
	shCmd := exec.Command("go", "get", "github.com/nimblehq/gin-templates")
	shCmd.Stdout = os.Stdout
	err := shCmd.Run()
	if err != nil {
		Fail("Failed to get gin-templates: " + err.Error())
	} else {
		log.Println("Gin-templates downloaded successfully.")
	}
}

func BuildGinTemplate() {
	shCmd := exec.Command("go", "build", "-o", os.Getenv("GOPATH")+"/bin/nimble-gin", "github.com/nimblehq/gin-templates")
	shCmd.Stdout = os.Stdout
	err := shCmd.Run()
	if err != nil {
		Fail("Failed to build gin-templates: " + err.Error())
	} else {
		log.Println("Gin-templates built successfully.")
	}
}

func RemoveCookiecuttersCache() {
	shCmd := exec.Command("rm", "-rf", os.Getenv("HOME")+"/.cookiecutters/gin-templates")
	shCmd.Stdout = os.Stdout

	err := shCmd.Run()
	if err != nil {
		Fail("Failed to remove cookiecutters cache: " + err.Error())
	} else {
		log.Println("Cookiecutters cache removed successfully.")
	}
}

func CreateProjectFromGinTemplate(input string) string {
	shCmd := exec.Command("cookiecutter", "../")

	stdin, err := shCmd.StdinPipe()
	if err != nil {
		Fail("Failed to get stdin pipe: " + err.Error())
	}

	go func() {
		defer stdin.Close()
		_, err = io.WriteString(stdin, input)
		if err != nil {
			Fail("Failed to write std value to file: " + err.Error())
		}
	}()

	output, err := shCmd.CombinedOutput()
	if err != nil {
		Fail("Failed to create template: " + err.Error() + "\n" + string(output))
	} else {
		log.Println("Template created successfully.")
	}

	currentPath, err := os.Getwd()
	if err != nil {
		Fail("Failed to get current directory: " + err.Error())
	}

	return currentPath
}
