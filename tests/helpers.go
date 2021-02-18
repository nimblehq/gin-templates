package tests

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
)

func ChangeDirectory(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		Fail("Failed to change directory: " + err.Error())
	}
}

func CreateTemplate() string {
	input := []byte("test-gin-templates")

	reader, writer, err := os.Pipe()
	if err != nil {
		Fail("Failed to connected pair of files: " + err.Error())
	}

	_, err = writer.Write(input)
	if err != nil {
		Fail("Failed to write file: " + err.Error())
	}
	writer.Close()

	shCmd := exec.Command("cookiecutter", "../..")
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

func ReadFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		Fail("Failed to read file: " + err.Error())
	}

	return string(content)
}
