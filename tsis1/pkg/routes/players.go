package routes

import (
	"encoding/json"
	"net/http"

	"github.com/margulan-kalykul/Golang/pkg/data"
	"github.com/margulan-kalykul/Golang/pkg/types"
)

func Players(w http.ResponseWriter, r *http.Request) {
	//declare response variable
	var response types.ResponsePlayers

	//Retrieve player details
	data.PrepareResponse()
	players := data.PrepareResponsePlayers()

	//assign player details to response
	response.Players = players

	//update content type
	w.Header().Set("Content-Type", "application/json")

	//specify HTTP status code
	w.WriteHeader(http.StatusOK)

	//convert struct to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	//update response
	w.Write(jsonResponse)
}