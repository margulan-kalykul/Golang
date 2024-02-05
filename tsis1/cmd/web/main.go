package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const PORT = ":8080"

func main(){
	router := mux.NewRouter()

	router.HandleFunc("/health-check", healthCheck).Methods("GET")
	router.HandleFunc("/films", responseFilms).Methods("GET")
	router.HandleFunc("/film/{id:[0-9]+}", responseFilm).Methods("GET")
	http.Handle("/", router)

	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}