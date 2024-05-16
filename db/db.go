package db

import (
	"database/sql"
	"fmt"

	"github.com/dragoneena12/lapi-hotel-system/config"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func NewDBConnection(config config.Config) (*sql.DB, error) {
	db, err := sql.Open(config.SQLDriver, config.DbName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	return db, nil
}
