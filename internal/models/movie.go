package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title    string `json:"title"`
	Year     string `json:"year"`
	Director string `json:"director"`
	Genre    string `json:"genre"`
	Country  string `json:"country"`
}
