package views

import (
	"net/http"
	"io"
	"github.com/gorilla/mux"
	"github.com/server-forecaster/model/manager"
	"encoding/json"
)

func GetByAlias(writer http.ResponseWriter, request *http.Request) {
	parameters := mux.Vars(request)
	user := manager.Create().GetUserByAlias(parameters["alias"])
	json.NewEncoder(writer).Encode(user)
	writer.WriteHeader(http.StatusAccepted)
}

func Insert(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusCreated)
	io.WriteString(writer, "SHUU!")
}
