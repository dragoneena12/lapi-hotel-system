package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/dragoneena12/lapi-hotel-system/domain"
)

const (
	TableNameHotels = "hotels"
)

type HotelRepository struct {
	db *sql.DB
}

func NewHotelRepository(db *sql.DB) *HotelRepository {
	return &HotelRepository{db}
}

func (r *HotelRepository) Create(hotel domain.Hotel) error {
	cmd := fmt.Sprintf("INSERT INTO %s (id, owner_id, name, location, carbon_awards, fullerene_awards, carbon_nanotube_awards, graphene_awards, diamond_awards, key) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", TableNameHotels)
	_, err := r.db.Exec(cmd, hotel.ID, hotel.OwnerID, hotel.Name, hotel.Location, strings.Join(hotel.CarbonAwards[:], ","), strings.Join(hotel.FullereneAwards[:], ","), strings.Join(hotel.CarbonNanotubeAwards[:], ","), strings.Join(hotel.GrapheneAwards[:], ","), strings.Join(hotel.DiamondAwards[:], ","), hotel.Key)
	if err != nil {
		return err
	}
	return err
}

func (r *HotelRepository) Update(hotel domain.Hotel) error {
	cmd := fmt.Sprintf("UPDATE %s SET owner_id = ?, name = ?, location = ?, carbon_awards = ?, fullerene_awards = ?, carbon_nanotube_awards = ?, graphene_awards = ?, diamond_awards = ? WHERE id = ?", TableNameHotels)
	_, err := r.db.Exec(cmd, hotel.OwnerID, hotel.Name, hotel.Location, strings.Join(hotel.CarbonAwards[:], ","), strings.Join(hotel.FullereneAwards[:], ","), strings.Join(hotel.CarbonNanotubeAwards[:], ","), strings.Join(hotel.GrapheneAwards[:], ","), strings.Join(hotel.DiamondAwards[:], ","), hotel.ID)
	if err != nil {
		return err
	}
	return err
}

func (r *HotelRepository) List(limit int) ([]*domain.Hotel, error) {
	cmd := fmt.Sprintf(`SELECT * FROM %s ORDER BY id ASC LIMIT ?`, TableNameHotels)
	rows, err := r.db.Query(cmd, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get hotels: %w", err)
	}
	defer rows.Close()

	var hotels []*domain.Hotel
	for rows.Next() {
		var hotel domain.Hotel
		var ar = make([]string, 5)
		rows.Scan(&hotel.ID, &hotel.OwnerID, &hotel.Name, &hotel.Location, &ar[0], &ar[1], &ar[2], &ar[3], &ar[4])
		hotel.CarbonAwards = strings.Split(ar[0], ",")
		hotel.FullereneAwards = strings.Split(ar[1], ",")
		hotel.CarbonNanotubeAwards = strings.Split(ar[2], ",")
		hotel.GrapheneAwards = strings.Split(ar[3], ",")
		hotel.DiamondAwards = strings.Split(ar[4], ",")
		hotels = append(hotels, &hotel)
	}
	return hotels, nil
}

func (r *HotelRepository) GetById(id string) (*domain.Hotel, error) {
	cmd := fmt.Sprintf(`SELECT id, owner_id, name, location, carbon_awards, fullerene_awards, carbon_nanotube_awards, graphene_awards, diamond_awards FROM %s WHERE id = ?`, TableNameHotels)
	row := r.db.QueryRow(cmd, id)
	var ar = make([]string, 5)
	var hotel domain.Hotel
	err := row.Scan(&hotel.ID, &hotel.OwnerID, &hotel.Name, &hotel.Location, &ar[0], &ar[1], &ar[2], &ar[3], &ar[4])
	if err != nil {
		return nil, fmt.Errorf("failed to get hotel: %w", err)
	}
	hotel.CarbonAwards = strings.Split(ar[0], ",")
	hotel.FullereneAwards = strings.Split(ar[1], ",")
	hotel.CarbonNanotubeAwards = strings.Split(ar[2], ",")
	hotel.GrapheneAwards = strings.Split(ar[3], ",")
	hotel.DiamondAwards = strings.Split(ar[4], ",")
	return &hotel, nil
}

func (r *HotelRepository) GetKeyById(id string) (string, error) {
	cmd := fmt.Sprintf(`SELECT key FROM %s WHERE id = ?`, TableNameHotels)
	row := r.db.QueryRow(cmd, id)
	var key string
	err := row.Scan(&key)
	if err != nil {
		return "", fmt.Errorf("failed to get key: %w", err)
	}
	return key, nil
}
