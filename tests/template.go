package tests

import (
	"io"
	"os/exec"

	. "github.com/onsi/ginkgo"
)

func (c Cookiecutter) CreateProjectFromGinTemplate(currentTemplatePath string) {
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

	ChangeDirectory(currentTemplatePath + "/" + c.AppName)
}
