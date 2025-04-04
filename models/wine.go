package models

type Wine struct {
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	Origin string  `json:"origin"`
	Year   int     `json:"year"`
	Price  float64 `json:"price"`
}
