package main

import (
	"net/http"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/server-forecaster/views"
	"github.com/gorilla/mux"
	"log"
)

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

	log.Panic(http.ListenAndServe(":"+port, router))
}
