package template_test

import (
	"log"
	"os"
	"os/exec"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var ProjectGeneratedPath string

func TestTemplate(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllers Suite")
}

var _ = BeforeSuite(func() {
	input := []byte("test-gin-templates")
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(input)
	if err != nil {
		log.Fatal(err)
	}
	w.Close()

	stdin := r

	shCmd := exec.Command("cookiecutter", "../..")
	shCmd.Stdout = os.Stdout
	shCmd.Stdin = stdin

	err = shCmd.Run()

	if err != nil {
		log.Fatal("Gin project created unsuccessfully: ", err.Error())
	} else {
		log.Print("Gin project created successfully.")
	}

	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	ProjectGeneratedPath = currentPath
})

var _ = AfterSuite(func() {
	os.RemoveAll(ProjectGeneratedPath + "/test-gin-templates")
})
