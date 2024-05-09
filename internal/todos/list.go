package todos

import (
	"inodaf/todo/internal/config"
	"inodaf/todo/internal/pkg/database"
	"inodaf/todo/internal/pkg/models"
)

type result struct {
	Item  *models.Item
	Index int
}

func ListDoneItems() ([]result, error) {
	var filtered = make([]result, 0)

	for index, item := range database.GetItems(config.DatabasePath) {
		if item.DoneAt != "" {
			filtered = append(filtered, result{Item: &item, Index: index})
		}
	}

	return filtered, nil
}

func ListPendingItems() ([]result, error) {
	var filtered = make([]result, 0)

	for index, item := range database.GetItems(config.DatabasePath) {
		if item.DoneAt != "" {
			continue
		}
		filtered = append(filtered, result{Item: &item, Index: index})
	}

	return filtered, nil
}
