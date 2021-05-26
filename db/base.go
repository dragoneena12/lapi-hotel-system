package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/dragoneena12/lapi-hotel-system/config"
	_ "github.com/mattn/go-sqlite3"
)

const (
	tableNameTodos = "todos"
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
            id INT PRIMARY KEY NOT NULL,
            text STRING,
						done STRING,
						userid INT)`, tableNameTodos)
	DbConnection.Exec(cmd)
}
