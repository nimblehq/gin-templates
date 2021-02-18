package tests

import (
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
)

func ChangeDirectory(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		Fail("Failed to change directory: " + err.Error())
	}
}

func ReadFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		Fail("Failed to read file: " + err.Error())
	}

	return string(content)
}
