package tests

import (
	"io/ioutil"

	. "github.com/onsi/ginkgo"
)

func ReadFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		Fail("Failed to read file: " + err.Error())
	}

	return string(content)
}
