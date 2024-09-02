package api

import (
	"log"
	"net/http"

	"github.com/KainoaGardner/webMahjongCalc/internal"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api/v1", func(r chi.Router) {
		handHandler := internal.NewHandler()
		handHandler.RegisterRoutes(r)

	})

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, r)

}
