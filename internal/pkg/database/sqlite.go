package database

import (
	"database/sql"
	"inodaf/todo/internal/config"

	_ "github.com/mattn/go-sqlite3"
)

func NewSQLiteStore(params ...string) (*sql.DB, error) {
	var sourceName string

	if len(params) == 0 {
		if path, err := config.GetDatabasePath(); err != nil {
			return nil, err
		} else {
			sourceName = path
		}
	} else {
		sourceName = params[0]
	}

	db, err := sql.Open("sqlite3", sourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
