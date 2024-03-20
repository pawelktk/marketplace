package models

type Product struct {
	Id          int     `json:"id" form:"id"`
	Name        string  `json:"name" form:"name"`
	Description string  `json:"description" form:"description"`
	Stock       int     `json:"stock" form:"stock"`
	Price       float64 `json:"price" form:"price"`
}
