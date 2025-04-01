package converter

import (
	"time"

	"github.com/gmalheiro/playground-golang-clean-arch/internal/domain/entity"
	"github.com/gmalheiro/playground-golang-clean-arch/internal/product/dto"
)

func ToDTO(input entity.Product) dto.ProductOutputDto {
	return dto.ProductOutputDto{
		ID:          input.ID,
		Name:        input.Name,
		Price:       input.Price,
		CodeValue:   input.CodeValue,
		Quantity:    input.Quantity,
		Expiration:  formatDate(input.Expiration),
		IsPublished: input.IsPublished,
	}
}

func formatDate(t time.Time) string {
	return t.Format("01-02-2006")
}
