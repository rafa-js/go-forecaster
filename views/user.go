package views

import (
	"net/http"
	"io"
)

func GetByAlias(writer http.ResponseWriter, request *http.Request) {
	//parameters := mux.Vars(request)
	//user := manager.Create().GetUserByAlias(parameters["alias"])
	//json.NewEncoder(writer).Encode(user)
	writer.WriteHeader(http.StatusAccepted)
	io.WriteString(writer, "SHUU!")
}

func Insert(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusCreated)
	io.WriteString(writer, "SHUU!")
}
