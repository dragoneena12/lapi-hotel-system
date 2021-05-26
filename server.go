package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dragoneena12/lapi-hotel-system/config"
	"github.com/dragoneena12/lapi-hotel-system/graph"
	"github.com/dragoneena12/lapi-hotel-system/graph/generated"
	"github.com/dragoneena12/lapi-hotel-system/graph/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/jwt"
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

	c := generated.Config{Resolvers: &graph.Resolver{}}
	c.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (interface{}, error) {
		router.Use(jwtauth.Verifier(tokenAuth))
		token, claims, _ := jwtauth.FromContext(ctx)
		if token == nil || jwt.Validate(token) != nil {
			return nil, fmt.Errorf(http.StatusText(http.StatusUnauthorized))
		}
		roles, ok := claims["https://lapi.tokyo/claims/roles"].([]model.Role)
		roles = append(roles, "USER")
		if !ok {
			return nil, fmt.Errorf("role error")
		}
		for _, r := range roles {
			if role == r {
				return next(ctx)
			}
		}
		return nil, fmt.Errorf("access denied")

	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))
	router.Handle("/graphql", srv)
	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))

	log.Printf("connect to http://localhost:%d/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), router))
}
