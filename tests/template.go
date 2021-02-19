package tests

import (
	"log"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
)

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

func CreateProjectFromGinTemplate() string {
	input := []byte("test-gin-templates")

	reader, writer, err := os.Pipe()
	if err != nil {
		Fail("Failed to create file to store stdin value: " + err.Error())
	}

	_, err = writer.Write(input)
	if err != nil {
		Fail("Failed to write stdin value to file: " + err.Error())
	}
	writer.Close()

	var shCmd *exec.Cmd
	if os.Getenv("BRANCH") == "" {
		shCmd = exec.Command("nimble-gin", "create")
	} else {
		shCmd = exec.Command("nimble-gin", "create", "-b", os.Getenv("BRANCH"))
	}

	shCmd.Stdout = os.Stdout
	shCmd.Stdin = reader

	err = shCmd.Run()
	if err != nil {
		Fail("Failed to create template: " + err.Error())
	} else {
		log.Println("Template created successfully.")
	}

	currentPath, err := os.Getwd()
	if err != nil {
		Fail("Failed to get current directory: " + err.Error())
	}

	return currentPath
}

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
