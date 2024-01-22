package main

import (
	"encoding/json"

	"github.com/inodaf/todo/models"
	"github.com/inodaf/todo/utils"
)

func add(title string, description string) {
	// title := flag.String("t", "", "Title of the todo")
	// description := flag.String("d", "", "Description of the todo")

	// flag.Parse()

	// if len(*title) == 0 {
	// 	fmt.Println("Please provide a title")
	// 	return
	// }

	item := *models.NewItem()
	item.Title = title
	item.Description = description

	items := utils.GetItems(DatabasePath)
	items = append(items, item)

	data, err := json.Marshal(items)
	if err != nil {
		panic(err)
	}

	utils.WriteItems(DatabasePath, data)
}
