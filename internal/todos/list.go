package todos

import (
	"errors"
	"inodaf/todo/internal/pkg/database"
	"inodaf/todo/internal/pkg/models"
)

var (
	ErrListingItems = errors.New("list: unable to list items")
)

func ListDoneItems() ([]*models.Item, error) {
	var items []*models.Item

	rows, err := database.DB.Query("SELECT * FROM todos WHERE done_at != '' ORDER BY done_at DESC")
	if err != nil {
		return nil, ErrListingItems
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.Id, &item.Title, &item.Description, &item.CreatedAt, &item.UpdatedAt, &item.DoneAt)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	return items, nil
}

func ListPendingItems() ([]*models.Item, error) {
	var items []*models.Item

	rows, err := database.DB.Query("SELECT * FROM todos WHERE done_at = '' ORDER BY created_at DESC")
	if err != nil {
		return nil, ErrListingItems
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.Id, &item.Title, &item.Description, &item.CreatedAt, &item.UpdatedAt, &item.DoneAt)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	return items, nil
}
