package todos

import (
	"errors"
	"inodaf/todo/internal/pkg/database"
	"inodaf/todo/internal/pkg/models"
)

var (
	ErrNotFoundItemRemove = errors.New("remove: the item does not exists")
	ErrItemIsNotDone      = errors.New("remove: the item is not done, confirm removal")
)

func Remove(itemID int, force bool) error {
	var count int
	var item models.Item

	err := database.DB.QueryRow("SELECT count(id), done_at FROM todos WHERE id = ?", itemID).Scan(&count, &item.DoneAt)
	if err != nil || count == 0 {
		return ErrNotFoundItemRemove
	}

	if len(item.DoneAt) == 0 && !force {
		return ErrItemIsNotDone
	}

	_, err = database.DB.Exec("DELETE FROM todos WHERE id = ?")
	if err != nil {
		return err
	}

	return nil
}
