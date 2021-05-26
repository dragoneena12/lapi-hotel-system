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
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", config.Config.JWTSecret, nil)

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	// _, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
	// fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}

func main() {
	port := config.Config.Port

	router := chi.NewRouter()
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
		r.Handle("/query", srv)
		r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())
			w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))
		})
	})

	// Public routes
	router.Group(func(r chi.Router) {
		router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	})

	log.Printf("connect to http://localhost:%d/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), router))
}
