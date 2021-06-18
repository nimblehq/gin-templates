package tests

import (
	"fmt"
	"io"
	"os/exec"

	. "github.com/onsi/ginkgo"
)

// This should be consistent with cookiecutter.json
type Cookiecutter struct {
	AppName string
}

// String order MUST be consistent with cookiecutter.json
func (c Cookiecutter) structToString() string {
	return fmt.Sprintf("%v", c.AppName)
}

func (c Cookiecutter) CreateProjectFromGinTemplate(templateGeneratedPath string) {
	shCmd := exec.Command("cookiecutter", "../")

	stdin, err := shCmd.StdinPipe()
	if err != nil {
		Fail("Failed to get stdin pipe: " + err.Error())
	}

	go func() {
		defer stdin.Close()
		_, err = io.WriteString(stdin, c.structToString())
		if err != nil {
			Fail("Failed to write std value to file: " + err.Error())
		}
	}()

	output, err := shCmd.CombinedOutput()
	if err != nil {
		Fail("Failed to create template: " + err.Error() + "\n" + string(output))
	}

	ChangeDirectory(templateGeneratedPath + "/" + c.AppName)
}
