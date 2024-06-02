package model

import (
	"time"

	"github.com/dragoneena12/lapi-hotel-system/domain"
)

type Stay struct {
	ID          string    `json:"id"`
	HotelID     string    `json:"hotelID"`
	Hotel       *Hotel    `json:"hotel"`
	CheckinTime time.Time `json:"checkinTime"`
}

func NewStayModel(stay domain.Stay) *Stay {
	return &Stay{
		ID:          stay.ID,
		HotelID:     stay.HotelID,
		CheckinTime: stay.CheckinTime,
	}
}

func NewStayModels(stays []*domain.Stay) []*Stay {
	var models []*Stay
	for _, stay := range stays {
		models = append(models, NewStayModel(*stay))
	}
	return models
}
