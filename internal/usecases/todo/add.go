package todo

import (
	"encoding/json"
	"errors"

	"inodaf/todo/internal/config"
	"inodaf/todo/internal/models"
	"inodaf/todo/internal/utils"
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
	items := utils.GetItems(config.DatabasePath)

	data, err := json.Marshal(append(items, *item))
	if err != nil {
		return ErrJSONCreationFailedAdd
	}

	utils.WriteItems(config.DatabasePath, data)
	return nil
}
