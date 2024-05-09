package todos

import (
	"encoding/json"
	"errors"
	"inodaf/todo/internal/config"
	"inodaf/todo/internal/pkg/database"
)

var ErrNotFoundItemRemove = errors.New("remove: the item does not exists")
var ErrItemIsNotDone = errors.New("remove: the item is not done, confirm removal")

func Remove(itemID int, force bool) error {
	items := database.GetItems(config.DatabasePath)
	if itemID > len(items)-1 {
		return ErrNotFoundItemRemove
	}

	if len(items[itemID].DoneAt) == 0 && !force {
		return ErrItemIsNotDone
	}

	// Remove an item by the index.
	items = append(items[:itemID], items[itemID+1:]...)

	// Convert the struct into a JSON string.
	data, err := json.Marshal(items)
	if err != nil {
		return ErrJSONCreationFailedEdit
	}

	// Save the newly updated JSON file.
	database.WriteItems(config.DatabasePath, data)
	return nil
}
