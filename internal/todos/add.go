package todos

import (
	"errors"
	"inodaf/todo/internal/pkg/database"
	"inodaf/todo/internal/pkg/models"
	"log"
)

var ErrSaveNewItemFailedAdd = errors.New("add: could not save new item")

type AddInput struct {
	Description string
	Title       string
}

func Add(input AddInput) error {
	item, err := models.NewItem(input.Title)
	if err != nil {
		return err
	}

	item.Description = input.Description

	stmt, err := database.DB.Prepare("INSERT INTO todos(title, description, created_at) VALUES(?, ?, ?)")
	if err != nil {
		log.Println("statement prepare error: ", err.Error())
		return ErrSaveNewItemFailedAdd
	}

	_, err = stmt.Exec(item.Title, item.Description, item.CreatedAt)
	if err != nil {
		log.Println("statement execution error: ", err.Error())
		return ErrSaveNewItemFailedAdd
	}

	return nil
}
