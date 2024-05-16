package db

import (
	"testing"
	"time"

	"github.com/dragoneena12/lapi-hotel-system/domain"
	"github.com/google/go-cmp/cmp"
)

func TestStayRepository_CreateAndGet(t *testing.T) {
	db, err := createDBForTest()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := NewStayRepository(db)
	now := time.Now()
	stayID, err := domain.GenerateStayID()
	if err != nil {
		t.Fatal(err)
	}

	stay := domain.Stay{
		ID:          stayID,
		HotelID:     "hotel1",
		UserID:      "user1",
		CheckinTime: now,
	}
	err = repo.Create(stay)
	if err != nil {
		t.Fatal(err)
	}

	got, err := repo.GetById(stayID)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(stay, *got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestStayRepository_CreateAndCount(t *testing.T) {
	db, err := createDBForTest()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := NewStayRepository(db)
	now := time.Now()
	stayID, err := domain.GenerateStayID()
	if err != nil {
		t.Fatal(err)
	}

	stay := domain.Stay{
		ID:          stayID,
		HotelID:     "hotel1",
		UserID:      "user1",
		CheckinTime: now,
	}
	err = repo.Create(stay)
	if err != nil {
		t.Fatal(err)
	}

	got, err := repo.Count("user1")
	if err != nil {
		t.Fatal(err)
	}
	if got != 1 {
		t.Errorf("want: 1, got: %d", got)
	}
}
