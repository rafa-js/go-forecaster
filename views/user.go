package views

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/server-forecaster/model/manager"
	"github.com/server-forecaster/utils"
	"github.com/server-forecaster/model/entity"
	"encoding/json"
	"github.com/liip/sheriff"
	"github.com/server-forecaster/model"
)

func GetByAlias(writer http.ResponseWriter, request *http.Request) {
	defer model.GetDatabase().Close()
	parameters := mux.Vars(request)
	userManager := manager.CreateUserManager()
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
}

func Insert(writer http.ResponseWriter, request *http.Request) {
	defer model.GetDatabase().Close()
	user := entity.User{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&user)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		user.Password = utils.MD5(user.Password)
		userManager := manager.CreateUserManager()
		if userManager.AddUser(&user) {
			writer.WriteHeader(http.StatusCreated)
		} else {
			writer.WriteHeader(http.StatusBadRequest)
		}
	}
}
