package views

import (
	"net/http"
	"io"
	"github.com/gorilla/mux"
	"github.com/server-forecaster/model/manager"
	"encoding/json"
	"github.com/liip/sheriff"
)

func GetByAlias(writer http.ResponseWriter, request *http.Request) {
	parameters := mux.Vars(request)
	user := manager.Create().GetUserByAlias(parameters["alias"])
	o := &sheriff.Options{
		Groups: []string{"visible"},
	}
	data, err := sheriff.Marshal(o, user)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		a, _ := json.MarshalIndent(data, "", "  ")
		//json.NewEncoder(writer).Encode(data)
		writer.WriteHeader(http.StatusAccepted)
		writer.Write(a)
	}
}

func Insert(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusCreated)
	io.WriteString(writer, "ADDED!")
}
