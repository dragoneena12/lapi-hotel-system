package graph

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/dragoneena12/lapi-hotel-system/controller"
	"github.com/dragoneena12/lapi-hotel-system/graph/generated"
)

func NewHandler(authController controller.AuthController, stayController controller.StayController, hotelController controller.HotelController, hasRole controller.HasRole) http.Handler {
	c := generated.Config{
		Resolvers: &Resolver{
			authController:  authController,
			stayController:  stayController,
			hotelController: hotelController,
		},
		Directives: generated.DirectiveRoot(*newDirective(hasRole)),
	}
	return handler.NewDefaultServer(generated.NewExecutableSchema(c))
}
