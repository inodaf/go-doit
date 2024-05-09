package todo

import (
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

	err := Edit(EditInput{ItemID: itemID, Item: &items[itemID] })
	if err != nil {
		return nil, ErrJSONCreationFailedMarkDone
	}

	return &items[itemID], nil
}

func MarkAsUndone(itemID int) (*models.Item, error) {
	items := utils.GetItems(config.DatabasePath)
	if itemID > len(items) {
		return nil, ErrJSONCreationFailedMarkDone
	}

	items[itemID].MarkAsUndone()

	err := Edit(EditInput{ ItemID: itemID, Item: &items[itemID] })
	if err != nil {
		return nil, ErrJSONCreationFailedMarkDone
	}

	return &items[itemID], nil
}
