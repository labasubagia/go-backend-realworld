package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/labasubagia/go-backend-realworld/port"
)

type Server struct {
	router  *chi.Mux
	service port.Service
}

func (server *Server) setupRouter() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	server.router = router
}

func NewServer(service port.Service) port.Server {
	server := &Server{
		service: service,
	}
	server.setupRouter()
	return server
}

func (server *Server) Start() error {
	return http.ListenAndServe(":8080", server.router)
}
