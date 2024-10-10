package todos

import (
	"errors"
	"inodaf/todo/internal/pkg/database"
	"inodaf/todo/internal/pkg/models"
	"time"
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

	stmt, err := database.DB.Prepare("INSERT INTO todos(title, description, created_at, updated_at, done_at) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return ErrSaveNewItemFailedAdd
	}

	_, err = stmt.Exec(item.Title, item.Description, item.CreatedAt.Format(time.DateTime), item.UpdatedAt.Format(time.DateTime), item.DoneAt.Format(time.DateTime))
	if err != nil {
		return ErrSaveNewItemFailedAdd
	}

	return nil
}
