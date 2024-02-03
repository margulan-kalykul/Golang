package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/margulan-kalykul/Golang/pkg/data"
)

func Player(w http.ResponseWriter, r *http.Request) {
	data.PrepareResponse()
	var id string = mux.Vars(r)["id"]
	i, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	response := data.PrepareResponseProfile(i-1)
	// if error != 0 {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}