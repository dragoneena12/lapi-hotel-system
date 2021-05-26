package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/dragoneena12/lapi-hotel-system/graph/generated"
	"github.com/dragoneena12/lapi-hotel-system/graph/model"
	jwtauth "github.com/go-chi/jwtauth/v5"
)

func (r *mutationResolver) Checkin(ctx context.Context, input model.Check) (*model.Stay, error) {
	_, claims, _ := jwtauth.FromContext(ctx)
	user, ok := claims["sub"].(string)
	if !ok {
		return nil, nil
	}
	stay := &model.Stay{
		ID:       fmt.Sprintf("T%d", rand.Int()),
		HotelId:  input.HotelID,
		Checkin:  time.Now(),
		Checkout: time.Time{},
		User:     user,
	}
	err := stay.Create()
	if err != nil {
		return stay, err
	}
	return stay, nil
}

func (r *mutationResolver) Checkout(ctx context.Context, input model.Check) (*model.Stay, error) {
	_, claims, _ := jwtauth.FromContext(ctx)
	user, ok := claims["sub"].(string)
	if !ok {
		return nil, nil
	}
	stay, err := model.GetMostRecentStay(user)
	if err != nil {
		return stay, err
	}
	stay.Checkout = time.Now()
	err = stay.Save()
	if err != nil {
		return stay, err
	}
	return stay, nil
}

func (r *mutationResolver) AddHotel(ctx context.Context, input model.NewHotel) (*model.Hotel, error) {
	_, claims, _ := jwtauth.FromContext(ctx)
	user, ok := claims["sub"].(string)
	if !ok {
		return nil, nil
	}
	hotel := &model.Hotel{
		ID:       fmt.Sprintf("T%d", rand.Int()),
		Name:     input.Name,
		Location: input.Location,
		Owner:    user,
	}
	err := hotel.Create()
	if err != nil {
		return hotel, err
	}
	return hotel, nil
}

func (r *queryResolver) Stays(ctx context.Context) ([]*model.Stay, error) {
	_, claims, _ := jwtauth.FromContext(ctx)
	user, ok := claims["sub"].(string)
	if !ok {
		return nil, nil
	}
	stays, err := model.GetAllStay(user, 100)
	return stays, err
}

func (r *queryResolver) Hotels(ctx context.Context) ([]*model.Hotel, error) {
	hotels, err := model.GetAllHotel(100)
	return hotels, err
}

func (r *stayResolver) Hotel(ctx context.Context, obj *model.Stay) (*model.Hotel, error) {
	return &model.Hotel{ID: obj.HotelId, Name: "Hotel " + obj.HotelId}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Stay returns generated.StayResolver implementation.
func (r *Resolver) Stay() generated.StayResolver { return &stayResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type stayResolver struct{ *Resolver }
