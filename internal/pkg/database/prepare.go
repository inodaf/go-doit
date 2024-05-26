package database

import (
	"database/sql"
	"errors"
)

var schema string = `CREATE TABLE IF NOT EXISTS todos (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	description TEXT,
	created_at TEXT NOT NULL,
	updated_at TEXT,
	done_at TEXT
);`

func Prepare(db *sql.DB) error {
	_, err := db.Exec(schema)
	if err != nil {
		return errors.New("db_prepare: failed to setup DB with schema")
	}

	return nil
}
