package tests

import (
	"os"

	. "github.com/onsi/ginkgo"
)

func GetCurrentDirectory() string {
	currentPath, err := os.Getwd()
	if err != nil {
		Fail("Failed to get current directory: " + err.Error())
	}

	return currentPath
}

func ChangeDirectory(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		Fail("Failed to change directory: " + err.Error())
	}
}
