package tests

import (
	"fmt"
	"reflect"
	"strings"
)

type Choices string

const (
	Yes Choices = "1"
	No  Choices = "2"
)

// Field and order MUST be the same as cookiecutter.json
type Cookiecutter struct {
	AppName   string
	UseLogrus Choices
}

func (c *Cookiecutter) fillDefaultValue() {
	if len(c.AppName) == 0 {
		c.AppName = "test-gin-templates"
	}

	if c.UseLogrus != Yes && c.UseLogrus != No {
		c.UseLogrus = No
	}
}

func (c *Cookiecutter) structToString() string {
	var structString strings.Builder

	c.fillDefaultValue()

	v := reflect.ValueOf(*c)
	for i := 0; i < v.NumField(); i++ {
		structString.WriteString(fmt.Sprintf("%v\n", v.Field(i)))
	}

	return structString.String()
}
