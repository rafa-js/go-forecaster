package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"net/http"
	"strings"
	"net/http/httptest"
	"net/url"
	"io/ioutil"
	"fmt"
)

type TestParameters struct {
	Handler            func(http.ResponseWriter, *http.Request)
	Description        string
	Url                string
	Method             string
	RequestBody        string
	QueryParameters    map[string]string
	ExpectedStatusCode int
	ExpectedBody       string
}

func RunTest(t *testing.T, test TestParameters) {
	req, err := http.NewRequest(test.Method, test.Url, strings.NewReader(test.RequestBody))
	assert.NoError(t, err)

	param := make(url.Values)
	for key, value := range test.QueryParameters {
		param[key] = []string{value}
	}

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(test.Handler)
	handler.ServeHTTP(resp, req)
	processRequest(t, req, resp)

	assertStatusCode(t, resp, &test)
	assertBody(t, resp, &test)

}

func RunTests(t *testing.T, tests []TestParameters) {
	for _, test := range tests {
		RunTest(t, test)
	}
}

func processRequest(t *testing.T, req *http.Request, resp *httptest.ResponseRecorder) {
	if p, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Fail()
	} else if strings.Contains(string(p), "Error") {
		t.Errorf("header response shouldn't return error: %s", p)
	}
}

func assertStatusCode(t *testing.T, resp *httptest.ResponseRecorder, test *TestParameters) {
	if test.ExpectedStatusCode == resp.Code {
		assert.Equal(t, test.ExpectedStatusCode, resp.Code, test.Description)
	} else {
		message := fmt.Sprintf("%v: Expected status code %v but got %v",
			test.Description, test.ExpectedStatusCode, resp.Code)
		assert.Fail(t, message)
	}
}

func assertBody(t *testing.T, resp *httptest.ResponseRecorder, test *TestParameters) {
	body := resp.Body.String()
	if test.ExpectedBody == body {
		assert.Equal(t, test.ExpectedBody, resp.Body.String(), test.Description)
	} else {
		message := fmt.Sprintf("%v: Expected body '%v' but got '%v'",
			test.Description, test.ExpectedBody, body)
		assert.Fail(t, message)
	}
}
