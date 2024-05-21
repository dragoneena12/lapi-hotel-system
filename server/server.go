package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dragoneena12/lapi-hotel-system/config"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

type Server struct {
	port        string
	debug       bool
	handler     http.Handler
	middlewares []func(http.Handler) http.Handler
}

func NewServer(config config.Config, handler http.Handler, middlewares ...func(http.Handler) http.Handler) *Server {
	return &Server{port: config.Port, debug: config.Debug, handler: handler, middlewares: middlewares}
}

func (s *Server) Start() error {
	router := chi.NewRouter()

	var origins []string
	if s.debug {
		origins = []string{"http://localhost:8000", "http://localhost:4000"}
	} else {
		origins = []string{"https://www.lapi.tokyo"}
	}

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   origins,
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		Debug:            s.debug,
	}).Handler)
	for _, middleware := range s.middlewares {
		router.Use(middleware)
	}
	router.Handle("/graphql", s.handler)

	if s.debug {
		router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
		slog.Info(fmt.Sprintf("connect to http://localhost:%s/ for GraphQL playground", s.port))
	}

	slog.Info(fmt.Sprintf("server started on port %s", s.port))
	return http.ListenAndServe(fmt.Sprintf(":%s", s.port), router)
}
