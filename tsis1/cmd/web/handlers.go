package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/margulan-kalykul/Golang/pkg/gameshow"
)

func respondWithJson(w http.ResponseWriter, status int, data interface{}) {
	response, err := json.Marshal(data)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, status int, message string) {
	respondWithJson(w, status, map[string]string{"error": message})
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, http.StatusOK, map[string]string{"info": gameshow.Info})
}

func responseFilms(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, http.StatusOK, gameshow.Films)
}

func responseFilm(w http.ResponseWriter, r *http.Request) {
	var id string = mux.Vars(r)["id"]
	i, _ := strconv.Atoi(id)
	film, err := gameshow.GetFilm(i)

	if err != nil {
		fmt.Println(err.Error())
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	
	respondWithJson(w, http.StatusOK, film)
}
