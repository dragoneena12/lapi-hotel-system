package model

import (
	"fmt"
	"strings"

	"github.com/dragoneena12/lapi-hotel-system/db"
)

type Hotel struct {
	ID                   string   `json:"id"`
	Name                 string   `json:"name"`
	Location             string   `json:"location"`
	Owner                string   `json:"owner"`
	CarbonAwards         []string `json:"carbonAwards"`
	FullereneAwards      []string `json:"fullereneAwards"`
	CarbonNanotubeAwards []string `json:"carbonNanotubeAwards"`
	GrapheneAwards       []string `json:"grapheneAwards"`
	DiamondAwards        []string `json:"diamondAwards"`
	Key                  string   `json:"key"`
}

func (c *Hotel) Create() error {
	cmd := fmt.Sprintf("INSERT INTO %s (name, location, owner, carbonAwards, fullereneAwards, carbonNanotubeAwards, grapheneAwards, diamondAwards, hotelKey) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", db.TableNameHotels)
	_, err := db.DbConnection.Exec(cmd, c.Name, c.Location, c.Owner, strings.Join(c.CarbonAwards[:], ","), strings.Join(c.FullereneAwards[:], ","), strings.Join(c.CarbonNanotubeAwards[:], ","), strings.Join(c.GrapheneAwards[:], ","), strings.Join(c.DiamondAwards[:], ","), c.Key)
	if err != nil {
		return err
	}
	return err
}

func (c *Hotel) Save() error {
	cmd := fmt.Sprintf("UPDATE %s SET name = ?, location = ?, owner = ?, carbonAwards = ?, fullereneAwards = ?, carbonNanotubeAwards = ?, grapheneAwards = ?, diamondAwards = ?, hotelKey = ? WHERE id = ?", db.TableNameHotels)
	_, err := db.DbConnection.Exec(cmd, c.Name, c.Location, c.Owner, strings.Join(c.CarbonAwards[:], ","), strings.Join(c.FullereneAwards[:], ","), strings.Join(c.CarbonNanotubeAwards[:], ","), strings.Join(c.GrapheneAwards[:], ","), strings.Join(c.DiamondAwards[:], ","), c.Key, c.ID)
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
		var ar = make([]string, 5)
		rows.Scan(&Hotel.ID, &Hotel.Name, &Hotel.Location, &Hotel.Owner, &ar[0], &ar[1], &ar[2], &ar[3], &ar[4], &Hotel.Key)
		Hotel.CarbonAwards = strings.Split(ar[0], ",")
		Hotel.FullereneAwards = strings.Split(ar[1], ",")
		Hotel.CarbonNanotubeAwards = strings.Split(ar[2], ",")
		Hotel.GrapheneAwards = strings.Split(ar[3], ",")
		Hotel.DiamondAwards = strings.Split(ar[4], ",")
		Hotels = append(Hotels, &Hotel)
	}
	return Hotels, nil
}

func GetHotelById(id string) (*Hotel, error) {
	cmd := fmt.Sprintf(`SELECT * FROM %s WHERE id = ?`, db.TableNameHotels)
	row := db.DbConnection.QueryRow(cmd, id)
	var ar = make([]string, 5)
	var Hotel Hotel
	err := row.Scan(&Hotel.ID, &Hotel.Name, &Hotel.Location, &Hotel.Owner, &ar[0], &ar[1], &ar[2], &ar[3], &ar[4], &Hotel.Key)
	if err != nil {
		return nil, err
	}
	Hotel.CarbonAwards = strings.Split(ar[0], ",")
	Hotel.FullereneAwards = strings.Split(ar[1], ",")
	Hotel.CarbonNanotubeAwards = strings.Split(ar[2], ",")
	Hotel.GrapheneAwards = strings.Split(ar[3], ",")
	Hotel.DiamondAwards = strings.Split(ar[4], ",")
	return &Hotel, nil
}
