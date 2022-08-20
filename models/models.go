package models

//This is used by gorm to create the tables IF they don't exist (which isn't obvious and is weird.)

type Product struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}
