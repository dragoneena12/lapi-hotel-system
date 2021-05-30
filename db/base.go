package db

import (
	"database/sql"
	"log"

	"github.com/dragoneena12/lapi-hotel-system/config"
	_ "github.com/go-sql-driver/mysql"
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
}
