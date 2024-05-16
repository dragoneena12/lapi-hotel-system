package controller

import (
	"fmt"

	"github.com/dragoneena12/lapi-hotel-system/controller/repository"
	"github.com/dragoneena12/lapi-hotel-system/domain"
	"github.com/pquerna/otp/totp"
)

type HotelController interface {
	Add(hotel domain.Hotel) (*domain.Hotel, error)
	Edit(userID string, newHotel domain.Hotel) (*domain.Hotel, error)
	List(limit int) ([]*domain.Hotel, error)
	GetById(id string) (*domain.Hotel, error)
	GetKeyById(id, userID string) (string, error)
}

type hotelController struct {
	hotelRepository repository.HotelRepository
}

func NewHotelController(hotelRepository repository.HotelRepository) *hotelController {
	return &hotelController{
		hotelRepository,
	}
}

func (c *hotelController) Add(hotel domain.Hotel) (*domain.Hotel, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "lapi.tokyo",
		AccountName: hotel.Name,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to generate key: %w", err)
	}
	hotel.Key = key.URL()
	hotel.ID, err = domain.GenerateHotelID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate hotel id: %w", err)
	}
	err = c.hotelRepository.Create(hotel)
	if err != nil {
		return nil, fmt.Errorf("failed to create hotel: %w", err)
	}
	return &hotel, nil
}

func (c *hotelController) Edit(userID string, newHotel domain.Hotel) (*domain.Hotel, error) {
	hotel, err := c.hotelRepository.GetById(newHotel.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get hotel: %w", err)
	}
	if hotel.OwnerID != userID {
		return nil, fmt.Errorf("you are not owner")
	}
	newHotel.Key = hotel.Key
	newHotel.OwnerID = hotel.OwnerID
	err = c.hotelRepository.Update(newHotel)
	if err != nil {
		return nil, fmt.Errorf("failed to update hotel: %w", err)
	}
	return &newHotel, nil
}

func (c *hotelController) List(limit int) ([]*domain.Hotel, error) {
	hotels, err := c.hotelRepository.List(limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get hotels: %w", err)
	}
	return hotels, nil
}

func (c *hotelController) GetById(id string) (*domain.Hotel, error) {
	hotel, err := c.hotelRepository.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get hotel: %w", err)
	}
	return hotel, nil
}

func (c *hotelController) GetKeyById(id, userID string) (string, error) {
	hotel, err := c.hotelRepository.GetById(id)
	if err != nil {
		return "", fmt.Errorf("failed to get hotel: %w", err)
	}
	if hotel.OwnerID != userID {
		return "", fmt.Errorf("you are not owner")
	}
	key, err := c.hotelRepository.GetKeyById(id)
	if err != nil {
		return "", fmt.Errorf("failed to get key: %w", err)
	}
	return key, nil
}
