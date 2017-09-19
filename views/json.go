package views

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/server-forecaster/model/entity"
	"github.com/server-forecaster/model/manager"
)

func jsonErr(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
}

func GetAuthToken(request *http.Request) string {
	return strings.Replace(request.Header.Get("Authorization"), "Basic ", "", 1)
}

func GetUserByToken(token string) *entity.User {
	return manager.CreateUserManager().GetUserByToken(token)
}
