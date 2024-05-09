package todos

import (
	"encoding/json"
	"errors"

	"inodaf/todo/internal/config"
	"inodaf/todo/internal/pkg/database"
	"inodaf/todo/internal/pkg/models"
)

var ErrJSONCreationFailedAdd = errors.New("Add: Could not build JSON string")

type AddInput struct {
	Description string
	Title       string
}

func Add(input AddInput) error {
	item, err := models.NewItem(input.Title)
	if err != nil {
		return err
	}

	item.Description = input.Description
	items := database.GetItems(config.DatabasePath)

	data, err := json.Marshal(append(items, *item))
	if err != nil {
		return ErrJSONCreationFailedAdd
	}

	database.WriteItems(config.DatabasePath, data)
	return nil
}
