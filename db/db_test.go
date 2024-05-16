package db

import (
	"database/sql"
	"fmt"
)

func createDBForTest() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "../test.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	stayResetQuery := fmt.Sprintf("DELETE FROM %s", TableNameStays)
	_, err = db.Exec(stayResetQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to reset table stay: %w", err)
	}
	hotelResetQuery := fmt.Sprintf("DELETE FROM %s", TableNameHotels)
	_, err = db.Exec(hotelResetQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to reset table hotel: %w", err)
	}
	return db, nil
}
