package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/margulan-kalykul/Golang/pkg/routes"
)

func main(){
	//create a new router
	router := mux.NewRouter()

	//specify endpoints, handler functions and HTTP method
	router.HandleFunc("/health-check", routes.HealthCheck).Methods("GET")
	router.HandleFunc("/players", routes.Players).Methods("GET")
	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)
}