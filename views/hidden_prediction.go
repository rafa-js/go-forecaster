package views

import (
	"net/http"
	"github.com/server-forecaster/model/manager"
	"encoding/json"
	"github.com/server-forecaster/model/entity"
)

func AddHiddenPrediction(writer http.ResponseWriter, request *http.Request) {
	hiddenPrediction := entity.HiddenPrediction{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&hiddenPrediction)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		fromUser := GetUserByToken(GetAuthToken(request))
		hiddenPrediction.FromUser = *fromUser
		hiddenPredictionManager := manager.CreateHiddenPredictionManager()
		defer hiddenPredictionManager.Close()
		if hiddenPredictionManager.AddPrediction(&hiddenPrediction) {
			writer.WriteHeader(http.StatusCreated)
		} else {
			writer.WriteHeader(http.StatusBadRequest)
		}
	}
}
