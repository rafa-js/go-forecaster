package main

import (
	"net/http"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/server-forecaster/views"
	"github.com/gorilla/mux"
	"log"
	"github.com/server-forecaster/model"
)

func wrap(handler func(writer http.ResponseWriter, request *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		model.OpenDB()
		handler(w, r)
		defer model.DB.Close()
	})
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "80"
	}

	router := mux.NewRouter()
	api := router.PathPrefix("/api").Subrouter()

	userApi := api.PathPrefix("/users").Subrouter()
	userApi.HandleFunc("/{alias}", wrap(views.GetByAlias)).Methods("GET")
	userApi.HandleFunc("", wrap(views.Insert)).Methods("POST")

	authApi := api.PathPrefix("/auth").Subrouter()
	authApi.HandleFunc("/login", wrap(views.Login)).Methods("POST")

	hiddenPredictionApi := api.PathPrefix("/hiddenPredictions").Subrouter()
	hiddenPredictionApi.HandleFunc("", wrap(views.AddHiddenPrediction)).Methods("POST")
	hiddenPredictionApi.HandleFunc("/{id}", wrap(views.UpdateHiddenPrediction)).Methods("PUT")
	hiddenPredictionApi.HandleFunc("/reveal", wrap(views.RevealHiddenPrediction)).Methods("POST")

	classificationApi := api.PathPrefix("/classification").Subrouter()
	classificationApi.HandleFunc("", wrap(views.GetClassification)).Methods("GET")

	log.Panic(http.ListenAndServe(":"+port, router))
}
