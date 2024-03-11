package todo

import (
	"encoding/json"
	"errors"

	"inodaf/todo/internal/config"
	"inodaf/todo/internal/models"
	"inodaf/todo/internal/utils"
)

var ErrNotFoundItemEdit = errors.New("edit: the \"item\" does not exists")
var ErrJSONCreationFailedEdit = errors.New("edit: could not build the JSON string")

type EditInput struct {
	ItemID int
	Item   *models.Item
}

func Edit(input EditInput) error {
	items := utils.GetItems(config.DatabasePath)
	if input.ItemID > len(items) {
		return ErrNotFoundItemEdit
	}

	// Update the item in the store.
	items[input.ItemID] = *input.Item

	// Convert the struct into a JSON string.
	data, err := json.Marshal(items)
	if err != nil {
		return ErrJSONCreationFailedEdit
	}

	// Save the newly updated JSON file.
	utils.WriteItems(config.DatabasePath, data)
	return nil
}
