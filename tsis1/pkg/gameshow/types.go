package gameshow

type Film struct {
	Id		int			`json:"id"`
	Title	string		`json:"title"`
	Year	int			`json:"year"`
	Rating	float32		`json:"rating"`	
	Genres	[]string	`json:"genres"`
}