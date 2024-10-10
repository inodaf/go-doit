package todos

import (
	"errors"

	"inodaf/todo/internal/pkg/database"
	"inodaf/todo/internal/pkg/models"
)

var ErrInvalidViewID = errors.New("view: invalid id")

func View(itemID int) (*models.Item, error) {
	var item models.Item

	err := database.DB.QueryRow("SELECT id, title, description, created_at, updated_at, done_at FROM todos WHERE id = ?", itemID).Scan(&item.Id, &item.Title, &item.Description, &item.CreatedAt, &item.UpdatedAt, &item.DoneAt)
	if err != nil {
		return nil, ErrInvalidViewID
	}

	return &item, nil
}
