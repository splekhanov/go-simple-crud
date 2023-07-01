package model

type Movie struct {
	ID       int    `json:"ID" example:"1" format:"int64" swaggerignore:"true"`
	Title    string `json:"title" example:"The Thing" format:"string"`
	Year     string `json:"year"`
	Director string `json:"director"`
	Genre    string `json:"genre"`
	Country  string `json:"country"`
}
