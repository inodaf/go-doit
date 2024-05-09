package todos

import (
	"encoding/json"
	"errors"
	"time"

	"inodaf/todo/internal/config"
	"inodaf/todo/internal/pkg/database"
	"inodaf/todo/internal/pkg/models"
)

var ErrNotFoundItemEdit = errors.New("edit: the \"item\" does not exists")
var ErrJSONCreationFailedEdit = errors.New("edit: could not build the JSON string")

type EditInput struct {
	ItemID int
	Item   *models.Item
}

func Edit(input EditInput) error {
	items := database.GetItems(config.DatabasePath)
	if input.ItemID > len(items) {
		return ErrNotFoundItemEdit
	}

	// Update the item in the store.
	items[input.ItemID] = *input.Item
	items[input.ItemID].UpdatedAt = time.Now().Format(time.RFC822)

	// Convert the struct into a JSON string.
	data, err := json.Marshal(items)
	if err != nil {
		return ErrJSONCreationFailedEdit
	}

	// Save the newly updated JSON file.
	database.WriteItems(config.DatabasePath, data)
	return nil
}
