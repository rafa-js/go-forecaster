package util

import (
	"net/http"
	"strings"
)

func GetAuthToken(request *http.Request) string {
	return strings.Replace(request.Header.Get("Authorization"), "Basic ", "", 1)
}
