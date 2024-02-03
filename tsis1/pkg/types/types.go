package types

type ResponsePlayers struct {
	Players []Player `json:"players"`
}

type Player struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Score     int    `json:"score"`
}

type Profile struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Score     int    `json:"score"`
	JoinDate  string `json:"joindate"`
}