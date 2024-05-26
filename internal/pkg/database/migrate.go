package database

import (
	"database/sql"
	"encoding/json"
	"inodaf/todo/internal/config"
	"inodaf/todo/internal/pkg/models"
	"log"
	"os"
	"sync"
)

func getJSONItems(path string) []models.Item {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var items []models.Item
	var decoder *json.Decoder = json.NewDecoder(file)

	err = decoder.Decode(&items)
	if err != nil {
		panic(err)
	}

	return items
}


func Migrate(db *sql.DB) {
	log.Println("💿 Starting JSON to SQL migration")
	items := getJSONItems(config.DatabasePath)
	var wg sync.WaitGroup

	stmt, err := db.Prepare("INSERT INTO todos(title, description, created_at, updated_at, done_at) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatalln("could not prepare SQL statement")
	}

	for _, item := range items {
		wg.Add(1)
		go func(i models.Item) {
			defer wg.Done()
			_, err := stmt.Exec(i.Title, i.Description, i.CreatedAt, i.UpdatedAt, i.DoneAt)

			if err != nil {
				log.Println("could not migrate item: ", i.Title, err.Error())
				return
			}

			log.Printf("📀 Item: \"%s\" migrated. \n", i.Title)
		}(item)
	}

	wg.Wait()
}

type todo struct {
	Id string `sql:"id"`
	DoneAt string `sql:"done_at"`
}

func UpdateDateTime(db *sql.DB) {
	log.Println("💿 Starting TimeStamps normalization")

	var wg sync.WaitGroup
	var todos []todo;

	updateAllTodos, err := db.Prepare("UPDATE todos SET done_at = ? WHERE id = ?")
	if err != nil {
		log.Fatalln("could not prepare update statement for UpdateDateTime")
	}

	rows, err := db.Query("SELECT id, done_at FROM todos ORDER BY done_at ASC LIMIT 6")
	if err != nil {
		log.Fatalln("could not query todos")
	}
	defer rows.Close()

	for rows.Next() {
		var item todo
		err := rows.Scan(&item.Id, &item.DoneAt)
		if err != nil {
			log.Fatal("could not scan: ", err.Error())
		}
		todos = append(todos, item)
	}

	for _, item := range todos {
		wg.Add(1)

		go func(i todo) {
			defer wg.Done()

			i.DoneAt = ""

			res, err := updateAllTodos.Exec(i.DoneAt, i.Id)
			if err != nil {
				log.Fatalln("could not perform update: ", err.Error())
			}

			updateCount, err := res.RowsAffected()
			if err != nil {
				log.Fatalln("could not get affected rows: ", err.Error())
			}

			log.Println("📀 Updated ", updateCount, " items")
		}(item)
	}

	wg.Wait()
}
