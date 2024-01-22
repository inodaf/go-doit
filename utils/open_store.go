package utils

import (
	"encoding/json"
	"os"

	"github.com/inodaf/todo/models"
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
