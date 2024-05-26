package todos

import (
	"errors"
	"time"

	"inodaf/todo/internal/pkg/database"
	"inodaf/todo/internal/pkg/models"
)

var (
	ErrNotFoundItemEdit       = errors.New("edit: item does not exists")
	ErrJSONCreationFailedEdit = errors.New("edit: could not build the JSON string")
)

type EditInput struct {
	Item   *models.Item
}

func Edit(input EditInput) error {
	var count int

	err := database.DB.QueryRow("SELECT count(id) FROM todos WHERE id = ?", input.Item.Id).Scan(&count)
	if err != nil || count == 0 {
		return ErrNotFoundItemEdit
	}

	input.Item.UpdatedAt = time.Now().Format(time.DateTime)

	_, err = database.DB.Exec("UPDATE todos SET title = ?, description = ?, updated_at = ?, done_at = ? WHERE id = ?", input.Item.Title, input.Item.Description, input.Item.UpdatedAt, input.Item.DoneAt, input.Item.Id)
	if err != nil {
		return err
	}

	return nil
}
