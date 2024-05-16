package db

import (
	"testing"

	"github.com/dragoneena12/lapi-hotel-system/domain"
	"github.com/google/go-cmp/cmp"
)

func TestHotelRepository_CreateAndGet(t *testing.T) {
	db, err := createDBForTest()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := NewHotelRepository(db)
	hotelID, err := domain.GenerateHotelID()
	if err != nil {
		t.Fatal(err)
	}

	hotel := domain.Hotel{
		ID:                   hotelID,
		OwnerID:              "owner1",
		Name:                 "hotel1",
		Location:             "location1",
		CarbonAwards:         []string{"award1"},
		FullereneAwards:      []string{"award2"},
		CarbonNanotubeAwards: []string{"award3"},
		GrapheneAwards:       []string{"award4"},
		DiamondAwards:        []string{"award5"},
		Key:                  "key1",
	}
	err = repo.Create(hotel)
	if err != nil {
		t.Fatal(err)
	}

	got, err := repo.GetById(hotelID)
	if err != nil {
		t.Fatal(err)
	}
	hotel.Key = "" // key is not returned by GetById
	if diff := cmp.Diff(hotel, *got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestHotelRepository_CreateAndGetKey(t *testing.T) {
	db, err := createDBForTest()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := NewHotelRepository(db)
	hotelID, err := domain.GenerateHotelID()
	if err != nil {
		t.Fatal(err)
	}

	hotel := domain.Hotel{
		ID:                   hotelID,
		OwnerID:              "owner1",
		Name:                 "hotel1",
		Location:             "location1",
		CarbonAwards:         []string{"award1"},
		FullereneAwards:      []string{"award2"},
		CarbonNanotubeAwards: []string{"award3"},
		GrapheneAwards:       []string{"award4"},
		DiamondAwards:        []string{"award5"},
		Key:                  "key1",
	}
	err = repo.Create(hotel)
	if err != nil {
		t.Fatal(err)
	}

	got, err := repo.GetKeyById(hotelID)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(hotel.Key, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}
