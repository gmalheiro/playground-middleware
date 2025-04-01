package dto

type CreateProductDto struct {
	Name        string  `json:"name,omitempty" validate:"required"`
	Quantity    int     `json:"quantity,omitempty" validate:"required"`
	CodeValue   string  `json:"code_value,omitempty" validate:"required"`
	IsPublished bool    `json:"is_published,omitempty" validate:"required"`
	Expiration  string  `json:"expiration,omitempty" validate:"required"`
	Price       float64 `json:"price,omitempty" validate:"required"`
}

type ProductOutputDto struct {
	ID          int     `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Quantity    int     `json:"quantity,omitempty"`
	CodeValue   string  `json:"code_value,omitempty"`
	IsPublished bool    `json:"is_published,omitempty"`
	Expiration  string  `json:"expiration,omitempty"`
	Price       float64 `json:"price,omitempty"`
}
