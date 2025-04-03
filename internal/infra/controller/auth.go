package controller

import (
	"github.com/gmalheiro/playground-golang-clean-arch/internal/infra/handler"
	"github.com/go-chi/chi/v5"
)

type AuthController struct {
	AuthHandler *handler.AuthHandler
}

func NewAuthController(authHandler *handler.AuthHandler) *AuthController {
	return &AuthController{
		AuthHandler: authHandler,
	}
}

func (ac *AuthController) RegisterRoutes(r chi.Router) {
	r.Route("/v1/auth", func(r chi.Router) {
		r.Post("/login", ac.AuthHandler.Login)
	})
}
