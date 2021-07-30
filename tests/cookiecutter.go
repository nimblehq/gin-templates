package tests

import (
	"fmt"
	"reflect"
	"strings"
)

type Choices string
type Variants string

const (
	Yes Choices = "1"
	No  Choices = "2"

	API  Variants = "1"
	Web  Variants = "2"
	Both Variants = "3"
)

// Field and order MUST be the same as cookiecutter.json
type Cookiecutter struct {
	AppName   string
	Variant   Variants
	UseLogrus Choices
	UseHeroku Choices
}

func (c *Cookiecutter) fillDefaultValue() {
	if len(c.AppName) == 0 {
		c.AppName = "test-gin-templates"
	}

	if c.UseLogrus != Yes && c.UseLogrus != No {
		c.UseLogrus = No
	}

	if c.UseHeroku != Yes && c.UseHeroku != No {
		c.UseHeroku = No
	}

	if c.Variant != API && c.Variant != Web && c.Variant != Both {
		c.Variant = API
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
