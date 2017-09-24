package views

import (
	"net/http"
	"github.com/server-forecaster/model/manager"
	"encoding/json"
	"github.com/liip/sheriff"
)

func GetClassification(writer http.ResponseWriter, request *http.Request) {
	classificationManager := manager.CreateClassificationManager()
	classification := classificationManager.GetClassification()
	o := &sheriff.Options{
		Groups: []string{"visible"},
	}
	data, err := sheriff.Marshal(o, classification.Scores[1])
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		encodedData, _ := json.Marshal(data)
		writer.WriteHeader(http.StatusAccepted)
		writer.Write(encodedData)
	}
}
