package database

import (
	"database/sql"
	"inodaf/todo/internal/config"

	_ "github.com/mattn/go-sqlite3"
)

func NewSQLiteStore(params ...string) (*sql.DB, error) {
	var sourceName string

	// What happens if the executable program is outside of the repository?
	// Should we create the DB in a temp location in the Home directory?
	if len(params) == 0 {
		sourceName = config.DatabasePath
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
