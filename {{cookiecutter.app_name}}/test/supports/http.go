package supports

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/nimblehq/{{cookiecutter.app_name}}/test"

	"github.com/onsi/ginkgo"
)

func MakeRequest(method string, path string, body io.Reader) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		ginkgo.Fail("Failed to init request: " + err.Error())
	}

	resp := httptest.NewRecorder()

	test.Router.ServeHTTP(resp, req)

	return resp
}
