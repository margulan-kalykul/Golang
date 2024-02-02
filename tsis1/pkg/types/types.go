package types

type Response struct{
	Players []Player `json:"players"`
}

type Player struct{
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}