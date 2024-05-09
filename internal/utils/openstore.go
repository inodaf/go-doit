package utils

import (
	"encoding/json"
	"os"

	"inodaf/todo/internal/models"
)

func GetItems(path string) []models.Item {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var items []models.Item
	var decoder *json.Decoder = json.NewDecoder(file)

	err = decoder.Decode(&items)
	if err != nil {
		panic(err)
	}

	return items
}

func WriteItems(path string, data []byte) {
	os.WriteFile(path, data, 0644)
}
