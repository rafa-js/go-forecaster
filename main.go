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

func wrap(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		defer model.GetDatabase().Close()
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
	userApi.HandleFunc("/{alias}", views.GetByAlias).Methods("GET")
	userApi.HandleFunc("", views.Insert).Methods("POST")

	authApi := api.PathPrefix("/auth").Subrouter()
	authApi.HandleFunc("/login", views.Login).Methods("POST")

	hiddenPredictionApi := api.PathPrefix("/hiddenPredictions").Subrouter()
	hiddenPredictionApi.HandleFunc("", views.AddHiddenPrediction).Methods("POST")
	hiddenPredictionApi.HandleFunc("/{id}", views.UpdateHiddenPrediction).Methods("PUT")
	hiddenPredictionApi.HandleFunc("/reveal", views.RevealHiddenPrediction).Methods("POST")

	log.Panic(http.ListenAndServe(":"+port, router))
}
