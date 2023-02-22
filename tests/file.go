package tests

import (
	"os"

	. "github.com/onsi/ginkgo"
)

func ReadFile(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		Fail("Failed to read file: " + err.Error())
	}

	return string(content)
}
