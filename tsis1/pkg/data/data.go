package data

import (
	"errors"

	"github.com/margulan-kalykul/Golang/pkg/types"
)

var profiles []types.Profile

func PrepareResponse() {
	profiles = []types.Profile{}
	var profile types.Profile

	profile.Id = 1
	profile.FirstName = "Issac"
	profile.LastName = "Newton"
	profile.Username = "INewton1643"
	profile.Score = 4700
	profile.JoinDate = "22.05.1670"
	profiles = append(profiles, profile)

	profile.Id = 2
	profile.FirstName = "Albert"
	profile.LastName = "Einstein"
	profile.Username = "EnergyIsMass"
	profile.Score = 3500
	profile.JoinDate = "06.11.1908"
	profiles = append(profiles, profile)
}

func PrepareResponsePlayers() []types.Player {
	var players []types.Player

	for i := 0; i < len(profiles); i++ {
		var player types.Player
		
		player.Id = profiles[i].Id
		player.FirstName = profiles[i].FirstName
		player.LastName = profiles[i].LastName
		player.Score = profiles[i].Score
		players = append(players, player)
	}

	return players
}

func PrepareResponseProfile(id int) (types.Profile, error) {
	var err error = nil
	if id >= len(profiles) || id <= 0 {
		err = errors.New("Such id doesn't exist")
		return types.Profile{Score: len(profiles)}, err
	}
	return profiles[id], err
}