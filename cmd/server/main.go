package main

import (
	"github.com/gmalheiro/playground-golang-clean-arch/configs"
	"github.com/gmalheiro/playground-golang-clean-arch/internal/infra/controller"
	"github.com/gmalheiro/playground-golang-clean-arch/internal/infra/handler"
	"github.com/gmalheiro/playground-golang-clean-arch/internal/infra/http"
	"github.com/gmalheiro/playground-golang-clean-arch/internal/infra/repository"
)

func main() {
	cfg := configs.NewConfig(".env.development")

	hs := http.NewHttpServer(cfg.Server.Port).SetupDefault()
	productRepository := repository.NewProductMemoryRepository()

	productHandler := handler.NewProductHandler(productRepository)

	hs.RegisterController(controller.NewProductController(productHandler))

	hs.Run()
}
