package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Stay struct {
	ID          string
	HotelID     string
	UserID      string
	CheckinTime time.Time
}

func GenerateStayID() (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", fmt.Errorf("failed to generate stay id: %w", err)
	}
	return id.String(), nil
}
