package todo

import (
	"errors"

	"inodaf/todo/internal/config"
	"inodaf/todo/internal/models"
	"inodaf/todo/internal/utils"
)

var ErrInvalidViewID = errors.New("view: invalid id")

func View(itemID int) (*models.Item, error) {
	items := utils.GetItems(config.DatabasePath)
	if itemID > len(items) {
		return nil, ErrInvalidViewID
	}

	return &items[itemID], nil
}
