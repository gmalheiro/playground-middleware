package http

import (
	"fmt"
	"net/http"
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

func (hs *HttpServer) RegisterController(controller Controller) {
	hs.Controllers = append(hs.Controllers, controller)
}

func (hs *HttpServer) loadControllers() {
	hs.Router.Route("/api", func(r chi.Router) {
		for _, controller := range hs.Controllers {
			controller.RegisterRoutes(r)
		}
	})
}

func (hs *HttpServer) Run() {
	hs.loadControllers()
	serverPort := fmt.Sprintf(":%s", hs.ServerPort)

	if err := http.ListenAndServe(serverPort, hs.Router); err != nil {
		panic(err)
	}
}
