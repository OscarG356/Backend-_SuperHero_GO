package models

// Biography define la estructura para la biografía de un héroe
type Biography struct {
	FullName string	`json:"full-name"`
}

// PowerStats define la estructura para las estadísticas de poder de un héroe
type PowerStats struct {
	Intelligence int `json:"intelligence"`
	Strength     int `json:"strength"`
	Speed        int `json:"speed"`
	Durability   int `json:"durability"`
	Power        int `json:"power"`
	Combat       int `json:"combat"`
}

// Images define la estructura para las imágenes de un héroe
type Images struct {
	XS string `json:"xs"`
	SM string `json:"sm"`
	MD string `json:"md"`
	LG string `json:"lg"`
}

// Hero define la estructura para representar un héroe
type Hero struct {
	Name       string `json:"name"`
	Biography  Biography `json:"biography"`	
	PowerStats PowerStats `json:"powerstats"`
	Images     Images `json:"images"`
}