package model

import (
	"fmt"

	"github.com/dragoneena12/lapi-hotel-system/db"
)

func (c *Hotel) Create() error {
	cmd := fmt.Sprintf("INSERT INTO %s (id, name, location) VALUES (?, ?, ?)", db.TableNameHotels)
	_, err := db.DbConnection.Exec(cmd, c.ID, c.Name, c.Location)
	if err != nil {
		return err
	}
	return err
}

func GetAllHotel(limit int) (Hotels []*Hotel, err error) {
	cmd := fmt.Sprintf(`SELECT * FROM %s ORDER BY id ASC LIMIT ?`, db.TableNameHotels)
	rows, err := db.DbConnection.Query(cmd, limit)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var Hotel Hotel
		rows.Scan(&Hotel.ID, &Hotel.Name, &Hotel.Location)
		Hotels = append(Hotels, &Hotel)
	}
	return Hotels, nil
}
