package todo

import (
	"encoding/json"
	"errors"

	"inodaf/todo/internal/config"
	"inodaf/todo/internal/models"
	"inodaf/todo/internal/utils"
)

var ErrNotFoundItemMarkDone = errors.New("markdone: the item does not exists")
var ErrJSONCreationFailedMarkDone = errors.New("markdone: failed to create JSON")

func MarkAsDone(itemID int) (*models.Item, error) {
	items := utils.GetItems(config.DatabasePath)
	if itemID > len(items) {
		return nil, ErrNotFoundItemMarkDone
	}

	items[itemID].MarkAsDone()

	data, err := json.Marshal(items)
	if err != nil {
		return nil, ErrJSONCreationFailedMarkDone
	}

	utils.WriteItems(config.DatabasePath, data)
	return &items[itemID], nil
}

func MarkAsUndone(itemID int) (*models.Item, error) {
	items := utils.GetItems(config.DatabasePath)
	if itemID > len(items) {
		return nil, ErrJSONCreationFailedMarkDone
	}

	items[itemID].MarkAsUndone()

	data, err := json.Marshal(items)
	if err != nil {
		return nil, ErrJSONCreationFailedMarkDone
	}

	utils.WriteItems(config.DatabasePath, data)
	return &items[itemID], nil
}
