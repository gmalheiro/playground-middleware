package contract

import "github.com/gmalheiro/playground-golang-clean-arch/internal/domain/entity"

type ProductRepository interface {
	GetAll() ([]entity.Product, error)
	GetByID(id int) (*entity.Product, error)
	GetByCodeValue(codeValue string) (*entity.Product, error)
	Create(product entity.Product) (*entity.Product, error)
	Update(product entity.Product) error
	Delete(id int) error
}
