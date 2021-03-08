package tests

import (
	"os"

	. "github.com/onsi/ginkgo"
)

func ChangeDirectory(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		Fail("Failed to change directory: " + err.Error())
	}
}
