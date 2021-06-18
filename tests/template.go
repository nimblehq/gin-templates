package tests

import (
	"io"
	"log"
	"os/exec"

	. "github.com/onsi/ginkgo"
)

func CreateProjectFromGinTemplate(input string, templateGeneratedPath string) {
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

	ChangeDirectory(templateGeneratedPath + "/test-gin-templates")
}
