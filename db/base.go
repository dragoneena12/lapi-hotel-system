package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/dragoneena12/lapi-hotel-system/config"
	_ "github.com/mattn/go-sqlite3"
)

const (
	TableNameStays  = "stays"
	TableNameHotels = "hotels"
)

var DbConnection *sql.DB

func init() {
	var err error
	DbConnection, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmd := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s (
            id STRING PRIMARY KEY NOT NULL,
            hotel_id STRING,
						checkin DATETIME,
						checkout DATETIME,
						user STRING)`, TableNameStays)
	DbConnection.Exec(cmd)

	cmd = fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s (
            id STRING PRIMARY KEY NOT NULL,
            name STRING,
						location STRING)`, TableNameHotels)
	DbConnection.Exec(cmd)
}