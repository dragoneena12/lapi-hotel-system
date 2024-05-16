package model

import "github.com/dragoneena12/lapi-hotel-system/domain"

func NewHotelModel(hotel domain.Hotel) *Hotel {
	return &Hotel{
		ID:                   hotel.ID,
		OwnerID:              hotel.OwnerID,
		Name:                 hotel.Name,
		Location:             hotel.Location,
		CarbonAwards:         hotel.CarbonAwards,
		FullereneAwards:      hotel.FullereneAwards,
		CarbonNanotubeAwards: hotel.CarbonNanotubeAwards,
		GrapheneAwards:       hotel.GrapheneAwards,
		DiamondAwards:        hotel.DiamondAwards,
	}
}

func NewHotelModels(hotels []*domain.Hotel) []*Hotel {
	var models []*Hotel
	for _, hotel := range hotels {
		models = append(models, NewHotelModel(*hotel))
	}
	return models
}
