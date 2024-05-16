package repository

import "github.com/dragoneena12/lapi-hotel-system/domain"

type HotelRepository interface {
	Create(hotel domain.Hotel) error
	Update(hotel domain.Hotel) error
	List(limit int) ([]*domain.Hotel, error)
	GetById(id string) (*domain.Hotel, error)
	GetKeyById(id string) (string, error)
}
