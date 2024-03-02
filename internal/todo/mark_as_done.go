package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"inodaf/todo/internal/config"
	"inodaf/todo/utils"
)

func MarkAsDone() {
	if len(os.Args) <= 2 {
		fmt.Println("Mark as Done: Please specify the item ID\nExample: `$ todo done 12`.")
		return
	}

	itemID, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Mark as Done: Please specify a valid item ID.")
		return
	}

	items := utils.GetItems(config.DatabasePath)
	if itemID > len(items) {
		fmt.Println("View: The item does not exists.")
		return
	}

	items[itemID].MarkAsDone()

	data, err := json.Marshal(items)
	if err != nil {
		panic("Failed to form a JSON")
	}

	utils.WriteItems(config.DatabasePath, data)
	utils.PrintItem(&items[itemID], itemID, false)
}

func MarkAsUndone() {
	if len(os.Args) <= 2 {
		fmt.Println("Mark as Undone: Please specify the item ID\nExample: `$ todo undone 12`.")
		return
	}

	itemID, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Mark as Undone: Please specify a valid item ID.")
		return
	}

	items := utils.GetItems(config.DatabasePath)
	if itemID > len(items) {
		fmt.Println("View: The item does not exists.")
		return
	}

	items[itemID].MarkAsUndone()

	data, err := json.Marshal(items)
	if err != nil {
		panic("Failed to form a JSON")
	}

	utils.WriteItems(config.DatabasePath, data)
	utils.PrintItem(&items[itemID], itemID, false)
}
