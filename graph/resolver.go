package graph

import (
	"github.com/dragoneena12/lapi-hotel-system/controller"
)

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	authController  controller.AuthController
	stayController  controller.StayController
	hotelController controller.HotelController
}
