package http

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Controller interface {
	RegisterRoutes(r chi.Router)
}

type HttpServer struct {
	Router      chi.Router
	ServerPort  string
	Controllers []Controller
}

func NewHttpServer(serverPort string) *HttpServer {
	return &HttpServer{
		Router:      chi.NewRouter(),
		ServerPort:  serverPort,
		Controllers: make([]Controller, 0),
	}
}
func (hs *HttpServer) SetupDefault() *HttpServer {
	hs.Router.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.StripSlashes,
		middleware.Timeout(5*time.Second),
		middleware.Heartbeat("/ping"),
	)

	return hs
}
