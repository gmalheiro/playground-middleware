package repository

import (
	"log"

	"github.com/gmalheiro/playground-golang-clean-arch/configs"
	"github.com/gmalheiro/playground-golang-clean-arch/internal/domain/entity"
	"github.com/gmalheiro/playground-golang-clean-arch/pkg/json_loader"
)

type ProductMemoryRepository struct {
	products     map[int]*entity.Product
	codeValueSet map[string]bool
}

func NewProductMemoryRepository() *ProductMemoryRepository {
	products, err := json_loader.Read(configs.GetEnv("FILE_PATH", ""))
	if err != nil {
		log.Fatal(err.Error())
	}
	return &ProductMemoryRepository{
		products:     products.Products,
		codeValueSet: products.Set,
	}
}
