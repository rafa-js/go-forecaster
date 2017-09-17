package views

import (
	"net/http"
)

type Person struct {
}

func GetPendingForecasts(writer http.ResponseWriter, r *http.Request) {
	writer.WriteHeader(http.StatusNotFound)
	writer.Write([]byte("SHU!"))
}
