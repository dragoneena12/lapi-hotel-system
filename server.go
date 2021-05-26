package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dragoneena12/lapi-hotel-system/config"
	"github.com/dragoneena12/lapi-hotel-system/graph"
	"github.com/dragoneena12/lapi-hotel-system/graph/generated"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/rs/cors"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", config.Config.JWTSecret, nil)
}

func main() {
	port := config.Config.Port

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8000"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	// Protected routes
	router.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))

		// Handle valid / invalid tokens. In this example, we use
		// the provided authenticator middleware, but you can write your
		// own very easily, look at the Authenticator method in jwtauth.go
		// and tweak it, its not scary.
		// r.Use(jwtauth.Authenticator)

		srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
		r.Handle("/v1", srv)
	})

	// Public routes
	router.Group(func(r chi.Router) {
		r.Handle("/", playground.Handler("GraphQL playground", "/v1"))
	})

	log.Printf("connect to http://localhost:%d/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), router))
}
