package gameshow

import "errors"

var Info string = "Author: Kalykul Margulan\nFilms collection app\nThis apps shows the list of some films and some info abput them like rating and genres."

var Films []Film = []Film{
	{
		Id:     1,
		Title:  "Back to the Future",
		Year:   1985,
		Rating: 8.5,
		Genres: []string{"adventure", "comedy", "sci-fi"},
	},
	{
		Id:     2,
		Title:  "Star Wars: Episode 1 - Fantom menace",
		Year:   1999,
		Rating: 6.5,
		Genres: []string{"action", "adventure", "fantasy"},
	},
	{
		Id:     3,
		Title:  "The Lord of the Rings: The Fellowship of the Ring",
		Year:   2001,
		Rating: 8.9,
		Genres: []string{"action", "adventure", "drama"},
	},
	{
		Id:     4,
		Title:  "Terminator 2: Judgment Day",
		Year:   1991,
		Rating: 8.6,
		Genres: []string{"action", "sci-fi"},
	},
	{
		Id:     5,
		Title:  "Inception",
		Year:   2010,
		Rating: 8.8,
		Genres: []string{"action", "adventure", "sci-fi"},
	},
}

func GetFilm(id int) (*Film, error) {
	if id <= 0 || id > len(Films) {
		return nil, errors.New("No such film in the database")
	}
	return &Films[id-1], nil
}
