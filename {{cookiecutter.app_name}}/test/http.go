package test

import (
	"io"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
)

func MakeRequest(method string, path string, body io.Reader) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		Fail("Failed to make a request: " + err.Error())
	}

	resp := httptest.NewRecorder()

	Router.ServeHTTP(resp, req)

	return resp
}
