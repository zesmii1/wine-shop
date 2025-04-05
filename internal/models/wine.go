package models

type Wine struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Year        int     `json:"year"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type WineEdit struct {
	Name        string  `json:"name"`
	Year        int     `json:"year"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}
