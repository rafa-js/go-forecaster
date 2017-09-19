package views

import (
	"net/http"
	"github.com/server-forecaster/model/manager"
	"encoding/json"
	"github.com/server-forecaster/model/entity"
	"github.com/server-forecaster/model"
)

func AddHiddenPrediction(writer http.ResponseWriter, request *http.Request) {
	defer model.GetDatabase().Close()
	hiddenPrediction := entity.HiddenPrediction{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&hiddenPrediction)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		fromUser := GetUserByToken(GetAuthToken(request))
		if fromUser == nil {
			writer.WriteHeader(http.StatusUnauthorized)
		} else {
			hiddenPrediction.FromUser = *fromUser
			hiddenPredictionManager := manager.CreateHiddenPredictionManager()
			if hiddenPredictionManager.AddPrediction(&hiddenPrediction) {
				writer.WriteHeader(http.StatusCreated)
			} else {
				writer.WriteHeader(http.StatusBadRequest)
			}
		}
	}
}
