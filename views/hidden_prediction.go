package views

import (
	"net/http"
	"github.com/server-forecaster/model/manager"
	"encoding/json"
	"github.com/server-forecaster/model/entity"
	"github.com/server-forecaster/util"
	"github.com/gorilla/mux"
	"strconv"
	"io"
)

func AddHiddenPrediction(writer http.ResponseWriter, request *http.Request) {
	hiddenPredictionManager := manager.CreateHiddenPredictionManager()
	handleSaveHiddenPrediction(writer, request, http.StatusCreated, func(hiddenPred *entity.HiddenPrediction) error {
		return hiddenPredictionManager.InsertPrediction(hiddenPred)
	})
}

func UpdateHiddenPrediction(writer http.ResponseWriter, request *http.Request) {
	parameters := mux.Vars(request)
	id, ok := parameters["id"]
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
	}

	hiddenPredictionManager := manager.CreateHiddenPredictionManager()
	handleSaveHiddenPrediction(writer, request, http.StatusAccepted, func(hiddenPred *entity.HiddenPrediction) error {
		predId, _ := strconv.Atoi(id)
		return hiddenPredictionManager.UpdatePrediction(predId, hiddenPred)
	})
}

func handleSaveHiddenPrediction(writer http.ResponseWriter,
	request *http.Request, successCode int, handle func(hiddenPred *entity.HiddenPrediction) error) {

	hiddenPrediction := entity.HiddenPrediction{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&hiddenPrediction)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		fromUser := GetUserByToken(util.GetAuthToken(request))
		if fromUser == nil {
			writer.WriteHeader(http.StatusUnauthorized)
		} else {
			hiddenPrediction.FromUser = *fromUser
			err := handle(&hiddenPrediction)
			if err == nil {
				writer.WriteHeader(successCode)
			} else {
				writer.WriteHeader(http.StatusBadRequest)
				io.WriteString(writer, err.Error())
			}
		}
	}
}

type RevealPredictionRequest struct {
	MatchId uint     `json:"matchId"`
	Secret  string   `json:"secret"`
}

func RevealHiddenPrediction(writer http.ResponseWriter, request *http.Request) {
	revealRequest := RevealPredictionRequest{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&revealRequest)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		fromUser := GetUserByToken(util.GetAuthToken(request))
		if fromUser == nil {
			writer.WriteHeader(http.StatusUnauthorized)
		} else {
			hiddenPredictionManager := manager.CreateHiddenPredictionManager()
			err := hiddenPredictionManager.RevealPrediction(revealRequest.Secret, revealRequest.MatchId, fromUser.ID)
			if err == nil {
				writer.WriteHeader(http.StatusCreated)
			} else {
				writer.WriteHeader(http.StatusBadRequest)
				io.WriteString(writer, err.Error())
				panic(err)
			}
		}
	}
}
