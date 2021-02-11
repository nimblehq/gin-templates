package tests

import (
	"io"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/gomega"
)

func MakeRequest(method string, path string, body io.Reader) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, path, body)
	Expect(err).NotTo(HaveOccurred())

	resp := httptest.NewRecorder()

	Router.ServeHTTP(resp, req)

	return resp
}
