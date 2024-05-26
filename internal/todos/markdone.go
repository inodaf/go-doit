package todos

import (
	"errors"

	"inodaf/todo/internal/pkg/database"
	"inodaf/todo/internal/pkg/models"
)

var (
	ErrNotFoundItemMarkDone = errors.New("markdone: item does not exists")
	ErrSaveMarkedDone       = errors.New("markdone: failed to edit item")
)

func MarkAsDone(itemID int) (*models.Item, error) {
	var item models.Item
	err := database.DB.QueryRow("SELECT * FROM todos WHERE id = ?", itemID).Scan(&item.Id, &item.Title, &item.Description, &item.CreatedAt, &item.UpdatedAt, &item.DoneAt)
	if err != nil {
		return nil, ErrNotFoundItemMarkDone
	}

	item.MarkAsDone()

	err = Edit(EditInput{Item: &item})
	if err != nil {
		return nil, ErrSaveMarkedDone
	}

	return &item, nil
}

func MarkAsUndone(itemID int) (*models.Item, error) {
	var item models.Item
	err := database.DB.QueryRow("SELECT * FROM todos WHERE id = ?", itemID).Scan(&item.Id, &item.Title, &item.Description, &item.CreatedAt, &item.UpdatedAt, &item.DoneAt)
	if err != nil {
		return nil, ErrNotFoundItemMarkDone
	}

	item.MarkAsUndone()

	err = Edit(EditInput{Item: &item})
	if err != nil {
		return nil, ErrSaveMarkedDone
	}

	return &item, nil
}
