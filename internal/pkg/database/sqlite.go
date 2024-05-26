package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// What happens if the executable program is outside of the repository?
// Should we create the DB in a temp location in the Home directory?
const defaultSourceName = "./data/live.db"

func Connect(params ...string) (*sql.DB, error) {
	var sourceName string;

	if len(params) == 0 {
		sourceName = defaultSourceName
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

