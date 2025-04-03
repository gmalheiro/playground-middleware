package handler

import (
	"net/http"

	"github.com/gmalheiro/playground-golang-clean-arch/internal/domain/contract"
	"github.com/gmalheiro/playground-golang-clean-arch/internal/product/dto"
	"github.com/gmalheiro/playground-golang-clean-arch/internal/product/usecase"
	"github.com/gmalheiro/playground-golang-clean-arch/pkg/web/request"
	"github.com/gmalheiro/playground-golang-clean-arch/pkg/web/response"
)

type ProductHandler struct {
	ProductRepository contract.ProductRepository
}

func NewProductHandler(repository contract.ProductRepository) *ProductHandler {
	return &ProductHandler{
		ProductRepository: repository,
	}
}

// Create RegisterRoutes @Summary Create new product
// @Description Create new product
// @Tags product
// @Accept  json
// @Produce  json
// @Param product body dto.CreateProductDto true "product"
// @Success 200 {object} entity.Product
// @Router /product [post]
func (ph *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateProductDto
	v, err := request.From(r).ParseValidJSON(&input)
	if err != nil {
		response.To(w).BadErr(err.Error()).SendJSON()
		return
	}

	if len(v) > 0 {
		response.To(w).UnprocessableErr(v).SendJSON()
		return
	}

	createProduct := usecase.NewCreateProductUseCase(ph.ProductRepository)

	output, resterr := createProduct.Execute(input)
	if resterr != nil {
		response.To(w).Status(resterr.StatusCode).MessageErr(resterr.Message).SendJSON()
		return
	}

	response.To(w).Status(http.StatusCreated).Content(output).SendJSON()
}
