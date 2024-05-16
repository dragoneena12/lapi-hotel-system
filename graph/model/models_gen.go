// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Hotel struct {
	ID                   string   `json:"id"`
	OwnerID              string   `json:"ownerID"`
	Name                 string   `json:"name"`
	Location             string   `json:"location"`
	CarbonAwards         []string `json:"carbonAwards"`
	FullereneAwards      []string `json:"fullereneAwards"`
	CarbonNanotubeAwards []string `json:"carbonNanotubeAwards"`
	GrapheneAwards       []string `json:"grapheneAwards"`
	DiamondAwards        []string `json:"diamondAwards"`
}

type HotelKey struct {
	Key string `json:"key"`
}

type Mutation struct {
}

type Query struct {
}

type Stay struct {
	ID          string    `json:"id"`
	HotelID     string    `json:"hotelID"`
	UserID      string    `json:"userID"`
	CheckinTime time.Time `json:"checkinTime"`
}

type Check struct {
	HotelID string `json:"hotelID"`
	Otp     string `json:"otp"`
}

type EditHotel struct {
	ID                   string   `json:"id"`
	Name                 string   `json:"name"`
	Location             string   `json:"location"`
	CarbonAwards         []string `json:"carbonAwards"`
	FullereneAwards      []string `json:"fullereneAwards"`
	CarbonNanotubeAwards []string `json:"carbonNanotubeAwards"`
	GrapheneAwards       []string `json:"grapheneAwards"`
	DiamondAwards        []string `json:"diamondAwards"`
}

type NewHotel struct {
	Name                 string   `json:"name"`
	Location             string   `json:"location"`
	CarbonAwards         []string `json:"carbonAwards"`
	FullereneAwards      []string `json:"fullereneAwards"`
	CarbonNanotubeAwards []string `json:"carbonNanotubeAwards"`
	GrapheneAwards       []string `json:"grapheneAwards"`
	DiamondAwards        []string `json:"diamondAwards"`
}

type Role string

const (
	RoleAdmin   Role = "ADMIN"
	RolePartner Role = "PARTNER"
	RoleUser    Role = "USER"
)

var AllRole = []Role{
	RoleAdmin,
	RolePartner,
	RoleUser,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RolePartner, RoleUser:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
