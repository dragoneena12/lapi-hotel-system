package repository

import "github.com/dragoneena12/lapi-hotel-system/domain"

type StayRepository interface {
	Create(stay domain.Stay) error
	Update(stay domain.Stay) error
	List(userID string, limit int) ([]*domain.Stay, error)
	GetMostRecent(userID string) (*domain.Stay, error)
	GetById(id string) (*domain.Stay, error)
	Count(userID string) (int, error)
}
