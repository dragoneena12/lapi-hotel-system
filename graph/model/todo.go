package model

import (
	"fmt"

	"github.com/dragoneena12/lapi-hotel-system/db"
)

type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"userid"`
}

const tableName = "todos"

func (c *Todo) Create() error {
	cmd := fmt.Sprintf("INSERT INTO %s (id, text, done, userid) VALUES (?, ?, ?, ?)", tableName)
	_, err := db.DbConnection.Exec(cmd, c.ID, c.Text, c.Done, c.UserID)
	if err != nil {
		return err
	}
	return err
}

func (c *Todo) Save() error {
	cmd := fmt.Sprintf("UPDATE %s SET text = ?, done = ?, userid = ? WHERE id = ?", tableName)
	_, err := db.DbConnection.Exec(cmd, c.Text, c.Done, c.UserID, c.ID)
	if err != nil {
		return err
	}
	return err
}

func GetAllTodo(userId string, limit int) (todos []*Todo, err error) {
	cmd := fmt.Sprintf(`SELECT id, text, done, userid FROM %s ORDER BY id ASC LIMIT ?`, tableName)
	rows, err := db.DbConnection.Query(cmd, limit)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var Todo Todo
		rows.Scan(&Todo.ID, &Todo.Text, &Todo.Done, &Todo.UserID)
		todos = append(todos, &Todo)
	}
	return todos, nil
}
