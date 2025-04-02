package usecase

import (
	"time"

	"github.com/gmalheiro/playground-golang-clean-arch/internal/domain/contract"
	"github.com/gmalheiro/playground-golang-clean-arch/internal/domain/entity"
	"github.com/gmalheiro/playground-golang-clean-arch/internal/product/converter"
	"github.com/gmalheiro/playground-golang-clean-arch/internal/product/dto"
	"github.com/gmalheiro/playground-golang-clean-arch/pkg/rest_err"
)

type CreateProductUseCase struct {
	ProductRepository contract.ProductRepository
}

func NewCreateProductUseCase(repository contract.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{
		ProductRepository: repository,
	}
}

func (puc *CreateProductUseCase) Execute(input dto.CreateProductDto) (*dto.ProductOutputDto, *rest_err.RestErr) {
	// product, err := puc.ProductRepository.GetByCodeValue(input.CodeValue)
	// if err != nil {
	// 	return nil, rest_err.NewInternalError(err.Error())
	// }
	// if product != nil {
	// 	return nil, rest_err.NewConflict(product_error.ErrConflictProductCodeValueAlreadyExist)
	// }
	parsedDate, err := parseDate(input.Expiration)
	if err != nil {
		return nil, rest_err.NewInternalError(err.Error())
	}
	productEntity := entity.NewProduct(input.Name, input.Quantity, input.CodeValue, input.IsPublished, parsedDate, input.Price)
	product, err := puc.ProductRepository.Create(*productEntity)
	if err != nil {
		return nil, rest_err.NewInternalError(err.Error())
	}
	output := converter.ToDTO(*product)
	return &output, nil
}

func parseDate(dateStr string) (time.Time, error) {
	layout := "02/01/2006"
	parsedTime, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}
