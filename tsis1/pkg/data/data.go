package data

import "github.com/margulan-kalykul/Golang/pkg/types"

func PrepareResponse() []types.Player {
	var players []types.Player
  
	var player types.Player
	player.Id = 1
	player.FirstName = "Issac"
	player.LastName = "Newton"
	players = append(players, player)
  
	player.Id = 2
	player.FirstName = "Albert"
	player.LastName = "Einstein"
	players = append(players, player)
  
	return players
  }