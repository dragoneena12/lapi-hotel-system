package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/dragoneena12/lapi-hotel-system/graph/generated"
	"github.com/dragoneena12/lapi-hotel-system/graph/model"
	jwtauth "github.com/go-chi/jwtauth/v5"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func (r *mutationResolver) Checkin(ctx context.Context, input model.Check) (*model.Stay, error) {
	_, claims, _ := jwtauth.FromContext(ctx)
	user, ok := claims["sub"].(string)
	if !ok {
		return nil, nil
	}
	stay, err := model.GetMostRecentStay(user)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	} else {
		if stay.Checkout.IsZero() {
			return nil, fmt.Errorf("already staying")
		}
	}
	hotel, err := model.GetHotelById(input.HotelID)
	if err != nil {
		return nil, err
	}
	key, err := otp.NewKeyFromURL(hotel.Key)
	if err != nil {
		return nil, err
	}
	valid := totp.Validate(input.Otp, key.Secret())
	if !valid {
		return nil, fmt.Errorf("provided OTP is not correct")
	}
	stay = &model.Stay{
		HotelId:  input.HotelID,
		Checkin:  time.Now(),
		Checkout: time.Time{},
		User:     user,
	}
	err = stay.Create()
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
	if !stay.Checkout.IsZero() {
		return nil, fmt.Errorf("no stay")
	}
	hotel, err := model.GetHotelById(input.HotelID)
	if err != nil {
		return nil, err
	}
	if hotel.ID != input.HotelID {
		return nil, fmt.Errorf("tried to check out wrong hotel")
	}
	key, err := otp.NewKeyFromURL(hotel.Key)
	if err != nil {
		return nil, err
	}
	valid := totp.Validate(input.Otp, key.Secret())
	if !valid {
		return nil, fmt.Errorf("provided OTP is not correct")
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
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "lapi.tokyo",
		AccountName: input.Name,
	})
	if err != nil {
		return nil, err
	}

	hotel := &model.Hotel{
		Name:                 input.Name,
		Location:             input.Location,
		Owner:                user,
		CarbonAwards:         input.CarbonAwards,
		FullereneAwards:      input.FullereneAwards,
		CarbonNanotubeAwards: input.CarbonNanotubeAwards,
		GrapheneAwards:       input.GrapheneAwards,
		DiamondAwards:        input.DiamondAwards,
		Key:                  key.URL(),
	}
	err = hotel.Create()
	if err != nil {
		return hotel, err
	}
	return hotel, nil
}

func (r *mutationResolver) EditHotel(ctx context.Context, input model.EditHotel) (*model.Hotel, error) {
	_, claims, _ := jwtauth.FromContext(ctx)
	user, ok := claims["sub"].(string)
	if !ok {
		return nil, nil
	}
	hotel, err := model.GetHotelById(input.ID)
	if err != nil {
		return hotel, err
	}
	if hotel.Owner != user {
		return nil, fmt.Errorf("you are not owner")
	}
	newHotel := &model.Hotel{
		ID:                   input.ID,
		Name:                 input.Name,
		Location:             input.Location,
		Owner:                user,
		CarbonAwards:         input.CarbonAwards,
		FullereneAwards:      input.FullereneAwards,
		CarbonNanotubeAwards: input.CarbonNanotubeAwards,
		GrapheneAwards:       input.GrapheneAwards,
		DiamondAwards:        input.DiamondAwards,
		Key:                  hotel.Key,
	}
	err = newHotel.Save()
	if err != nil {
		return newHotel, err
	}
	return newHotel, nil
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

func (r *queryResolver) Hotel(ctx context.Context, id string) (*model.Hotel, error) {
	hotel, err := model.GetHotelById(id)
	if err != nil {
		return nil, err
	}

	return hotel, nil
}

func (r *queryResolver) HotelKey(ctx context.Context, id string) (*model.HotelKey, error) {
	_, claims, _ := jwtauth.FromContext(ctx)
	user, ok := claims["sub"].(string)
	if !ok {
		return nil, nil
	}
	hotel, err := model.GetHotelById(id)
	if err != nil {
		return nil, err
	}
	if hotel.Owner != user {
		return nil, fmt.Errorf("you are not owner")
	}
	hotelKey := &model.HotelKey{
		Key: hotel.Key,
	}
	return hotelKey, nil
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
