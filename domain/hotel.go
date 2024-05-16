package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type Hotel struct {
	ID                   string
	OwnerID              string
	Name                 string
	Location             string
	CarbonAwards         []string
	FullereneAwards      []string
	CarbonNanotubeAwards []string
	GrapheneAwards       []string
	DiamondAwards        []string
	Key                  string
}

func GenerateHotelID() (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", fmt.Errorf("failed to generate hotel id: %w", err)
	}
	return id.String(), nil
}
