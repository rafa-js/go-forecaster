package views

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/server-forecaster/model/manager"
	"encoding/json"
	"github.com/liip/sheriff"
	"github.com/server-forecaster/model/entity"
)

func GetByAlias(writer http.ResponseWriter, request *http.Request) {
	parameters := mux.Vars(request)
	userManager := manager.Create()
	user := userManager.GetUserByAlias(parameters["alias"])
	if user == nil {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		o := &sheriff.Options{
			Groups: []string{"visible"},
		}
		data, err := sheriff.Marshal(o, user)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		} else {
			encodedData, _ := json.Marshal(data)
			writer.WriteHeader(http.StatusAccepted)
			writer.Write(encodedData)
		}
	}
	userManager.Close()
}

func Insert(writer http.ResponseWriter, request *http.Request) {
	user := entity.User{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&user)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		userManager := manager.Create()
		if userManager.AddUser(&user) {
			writer.WriteHeader(http.StatusCreated)
		} else {
			writer.WriteHeader(http.StatusBadRequest)
		}
		userManager.Close()
	}
}
