package routes

import (
	"fmt"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	//specify status code
	w.WriteHeader(http.StatusOK)

	//update response writer
	fmt.Fprintf(w, "Author: Kalykul Margulan\nPlayers leaderboard app\nThis apps shows the current board of players and can return the profile of the individual player")
}