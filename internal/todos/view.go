package todos

import (
	"errors"

	"inodaf/todo/internal/config"
	"inodaf/todo/internal/pkg/database"
	"inodaf/todo/internal/pkg/models"
)

var ErrInvalidViewID = errors.New("view: invalid id")

func View(itemID int) (*models.Item, error) {
	items := database.GetItems(config.DatabasePath)
	if itemID > len(items) {
		return nil, ErrInvalidViewID
	}

	return &items[itemID], nil
}
