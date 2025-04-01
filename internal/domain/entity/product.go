package entity

import "time"

type Product struct {
	ID          int
	Name        string
	Quantity    int
	CodeValue   string
	IsPublished bool
	Expiration  time.Time
	Price       float64
}

func NewProduct(name string, quantity int, codeValue string, isPublished bool, expiration time.Time, price float64) *Product {
	return &Product{
		Name:        name,
		Quantity:    quantity,
		CodeValue:   codeValue,
		IsPublished: isPublished,
		Expiration:  expiration,
		Price:       price,
	}
}
