package model

import "github.com/dragoneena12/lapi-hotel-system/domain"

func NewStayModel(stay domain.Stay) *Stay {
	return &Stay{
		ID:          stay.ID,
		HotelID:     stay.HotelID,
		UserID:      stay.UserID,
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
