package controller

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/dragoneena12/lapi-hotel-system/controller/repository"
	"github.com/dragoneena12/lapi-hotel-system/domain"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type StayController interface {
	Checkin(stay domain.Stay, passcode string) (*domain.Stay, error)
	List(userID string, limit int) ([]*domain.Stay, error)
	Count(userID string) (int, error)
}

type stayController struct {
	stayRepository  repository.StayRepository
	hotelRepository repository.HotelRepository
}

func NewStayController(stayRepository repository.StayRepository, hotelRepository repository.HotelRepository) *stayController {
	return &stayController{
		stayRepository,
		hotelRepository,
	}
}

func (c *stayController) Checkin(stay domain.Stay, passcode string) (*domain.Stay, error) {
	now := time.Now()
	count, err := c.stayRepository.Count(stay.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get stay count: %w", err)
	}
	if count > 0 {
		recentStay, err := c.stayRepository.GetMostRecent(stay.UserID)
		if err != nil {
			return nil, fmt.Errorf("failed to get most recent stay: %w", err)
		}
		if recentStay.CheckinTime.After(now.Add(-24 * time.Hour)) {
			return nil, fmt.Errorf("checkin is only allowed once per day")
		}
	}
	hotelKey, err := c.hotelRepository.GetKeyById(stay.HotelID)
	if err != nil {
		return nil, fmt.Errorf("failed to get hotel: %w", err)
	}
	key, err := otp.NewKeyFromURL(hotelKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get key: %w", err)
	}
	valid := totp.Validate(passcode, key.Secret())
	if !valid {
		return nil, fmt.Errorf("provided OTP is not correct")
	}
	stay.CheckinTime = now
	stay.ID, err = domain.GenerateStayID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate stay id: %w", err)
	}
	err = c.stayRepository.Create(stay)
	if err != nil {
		return nil, fmt.Errorf("failed to create stay: %w", err)
	}
	slog.Info("Checkin", "userID", stay.UserID, "hotelID", stay.HotelID, "time", stay.CheckinTime)
	return &stay, nil
}

func (c *stayController) List(userID string, limit int) ([]*domain.Stay, error) {
	stays, err := c.stayRepository.List(userID, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get stays: %w", err)
	}
	return stays, nil
}

func (c *stayController) Count(userID string) (int, error) {
	result, err := c.stayRepository.Count(userID)
	if err != nil {
		return 0, fmt.Errorf("failed to get stays: %w", err)
	}
	return result, nil
}
