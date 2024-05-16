package db

import (
	"database/sql"
	"fmt"

	"github.com/dragoneena12/lapi-hotel-system/domain"
)

const (
	TableNameStays = "stays"
)

type StayRepository struct {
	db *sql.DB
}

func NewStayRepository(db *sql.DB) *StayRepository {
	return &StayRepository{db}
}

func (r *StayRepository) Create(stay domain.Stay) error {
	cmd := fmt.Sprintf("INSERT INTO %s (id, user_id, hotel_id, checkin) VALUES (?, ?, ?, ?)", TableNameStays)
	_, err := r.db.Exec(cmd, stay.ID, stay.UserID, stay.HotelID, stay.CheckinTime)
	if err != nil {
		return err
	}
	return err
}

func (r *StayRepository) Update(stay domain.Stay) error {
	cmd := fmt.Sprintf("UPDATE %s SET user_id = ?, hotel_id = ?, checkin = ? WHERE id = ?", TableNameStays)
	_, err := r.db.Exec(cmd, stay.UserID, stay.HotelID, stay.CheckinTime, stay.ID)
	if err != nil {
		return err
	}
	return err
}

func (r *StayRepository) List(userID string, limit int) ([]*domain.Stay, error) {
	cmd := fmt.Sprintf(`SELECT * FROM %s WHERE user_id = ? ORDER BY id DESC LIMIT ?`, TableNameStays)
	rows, err := r.db.Query(cmd, userID, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get stays: %w", err)
	}
	defer rows.Close()

	var stays []*domain.Stay
	for rows.Next() {
		var stay domain.Stay
		rows.Scan(&stay.ID, &stay.UserID, &stay.HotelID, &stay.CheckinTime)
		stays = append(stays, &stay)
	}
	return stays, nil
}

func (r *StayRepository) GetMostRecent(userID string) (*domain.Stay, error) {
	cmd := fmt.Sprintf(`SELECT * FROM %s WHERE user_id = ? ORDER BY id DESC LIMIT 1`, TableNameStays)
	row := r.db.QueryRow(cmd, userID)
	var stay domain.Stay
	err := row.Scan(&stay.ID, &stay.UserID, &stay.HotelID, &stay.CheckinTime)
	if err != nil {
		return nil, err
	}
	return &stay, nil
}

func (r *StayRepository) GetById(id string) (*domain.Stay, error) {
	cmd := fmt.Sprintf(`SELECT * FROM %s WHERE id = ?`, TableNameStays)
	row := r.db.QueryRow(cmd, id)
	var stay domain.Stay
	err := row.Scan(&stay.ID, &stay.UserID, &stay.HotelID, &stay.CheckinTime)
	if err != nil {
		return nil, err
	}
	return &stay, nil
}

func (r *StayRepository) Count(userID string) (int, error) {
	cmd := fmt.Sprintf(`SELECT COUNT(*) FROM %s WHERE user_id = ?`, TableNameStays)
	row := r.db.QueryRow(cmd, userID)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
