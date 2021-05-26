package model

import (
	"fmt"
	"time"

	"github.com/dragoneena12/lapi-hotel-system/db"
)

type Stay struct {
	ID       string    `json:"id"`
	HotelId  string    `json:"hotel_id"`
	Checkin  time.Time `json:"checkin"`
	Checkout time.Time `json:"checkout"`
	User     string    `json:"user"`
}

func (c *Stay) Create() error {
	cmd := fmt.Sprintf("INSERT INTO %s (id, hotel_id, checkin, checkout, user) VALUES (?, ?, ?, ?, ?)", db.TableNameStays)
	_, err := db.DbConnection.Exec(cmd, c.ID, c.HotelId, c.Checkin, c.Checkout, c.User)
	if err != nil {
		return err
	}
	return err
}

func (c *Stay) Save() error {
	cmd := fmt.Sprintf("UPDATE %s SET hotel_id = ?, checkin = ?, checkout = ? user = ? WHERE id = ?", db.TableNameStays)
	_, err := db.DbConnection.Exec(cmd, c.HotelId, c.Checkin, c.Checkout, c.User, c.ID)
	if err != nil {
		return err
	}
	return err
}

func GetAllStay(user string, limit int) (Stays []*Stay, err error) {
	cmd := fmt.Sprintf(`SELECT * FROM %s WHERE user = ? ORDER BY id DESC LIMIT ?`, db.TableNameStays)
	rows, err := db.DbConnection.Query(cmd, user, limit)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var Stay Stay
		rows.Scan(&Stay.ID, &Stay.HotelId, &Stay.Checkin, &Stay.Checkout, &Stay.User)
		Stays = append(Stays, &Stay)
	}
	return Stays, nil
}

func GetMostRecentStay(user string) (Stay *Stay, err error) {
	cmd := fmt.Sprintf(`SELECT * FROM %s WHERE user = ? ORDER BY id DESC LIMIT 1`, db.TableNameStays)
	row := db.DbConnection.QueryRow(cmd, user)
	err = row.Scan(&Stay.ID, &Stay.HotelId, &Stay.Checkin, &Stay.Checkout, &Stay.User)
	if err != nil {
		return
	}
	return Stay, nil
}
