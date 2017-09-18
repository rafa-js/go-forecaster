package api

import (
	"github.com/server-forecaster/tests"
	"net/http"
	"github.com/server-forecaster/views"
)

func GetTests() []tests.TestParameters {
	return []tests.TestParameters{
		{
			Handler:            views.GetByAlias,
			Description:        "Get a not existent user",
			Url:                "api/users/notExistent",
			Method:             http.MethodGet,
			RequestBody:        ``,
			QueryParameters:    make(map[string]string),
			ExpectedStatusCode: 404,
			ExpectedBody:       ``,
		},
		{
			Handler:            views.GetByAlias,
			Description:        "Get an user with ID = 1",
			Url:                "api/users/alias",
			Method:             http.MethodGet,
			RequestBody:        ``,
			QueryParameters:    make(map[string]string),
			ExpectedStatusCode: 202,
			ExpectedBody:       `{"alias":"userAlias"}`,
		},
		{
			Handler:            views.Insert,
			Description:        "Add a new user",
			Url:                "api/users",
			Method:             http.MethodPost,
			RequestBody:        `{"alias":"userAlias", "password": "123", "email": "email@test.com"}`,
			QueryParameters:    make(map[string]string),
			ExpectedStatusCode: 201,
			ExpectedBody:       ``,
		},
	}
}
