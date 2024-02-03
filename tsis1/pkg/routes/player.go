package routes

import (
	"encoding/json"
	"fmt"
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
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response, err := data.PrepareResponseProfile(i-1)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	w.Write(jsonResponse)
}