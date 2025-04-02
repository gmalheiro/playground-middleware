package controller

import (
	"github.com/gmalheiro/playground-golang-clean-arch/internal/infra/handler"
	"github.com/go-chi/chi/v5"
)

type ProductController struct {
	ProductHandler *handler.ProductHandler
}

func NewProductController(productHandler *handler.ProductHandler) *ProductController {
	return &ProductController{
		ProductHandler: productHandler,
	}
}
func (pc *ProductController) RegisterRoutes(r chi.Router) {
	r.Route("/v1/products", func(r chi.Router) {
		r.Post("/", pc.ProductHandler.Create)
	})
}
