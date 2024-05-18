// models.go

package models

// esto datos se tomaron teniendo encuenta la solicitud de una contancia para el caso de la UdeA
type Biography struct {
	FullName string `json:"fullName"`
}

type Powerstats struct {
	Intelligence int `json:"intelligence"`
	Strength     int `json:"strength"`
	Speed        int `json:"speed"`
	Durability   int `json:"durability"`
	Power        int `json:"power"`
	Combat       int `json:"combat"`
}

type Images struct {
	XS string `json:"xs"`
	SM string `json:"sm"`
	MD string `json:"md"`
	LG string `json:"lg"`
}

type Superhero struct {
	Name       string     `json:"name"`
	Biography  Biography  `json:"biography"`
	Powerstats Powerstats `json:"powerstats"`
	Images     Images     `json:"images"`
}
