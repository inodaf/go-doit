package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/inodaf/todo/utils"
)

func markAsDone() {
	if len(os.Args) <= 2 {
		fmt.Println("Mark as Done: Please specify the item ID\nExample: `$ todo done 12`.")
		return
	}

	itemID, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Mark as Done: Please specify a valid item ID.")
		return
	}

	items := utils.GetItems(DatabasePath)
	if itemID > len(items) {
		fmt.Println("View: The item does not exists.")
		return
	}

	items[itemID].MarkAsDone()

	data, err := json.Marshal(items)
	if err != nil {
		panic("Failed to form a JSON")
	}

	utils.WriteItems(DatabasePath, data)
	utils.PrintItem(&items[itemID], itemID, false)
}

func markAsUndone() {
	if len(os.Args) <= 2 {
		fmt.Println("Mark as Undone: Please specify the item ID\nExample: `$ todo undone 12`.")
		return
	}

	itemID, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Mark as Undone: Please specify a valid item ID.")
		return
	}

	items := utils.GetItems(DatabasePath)
	if itemID > len(items) {
		fmt.Println("View: The item does not exists.")
		return
	}

	items[itemID].MarkAsUndone()

	data, err := json.Marshal(items)
	if err != nil {
		panic("Failed to form a JSON")
	}

	utils.WriteItems(DatabasePath, data)
	utils.PrintItem(&items[itemID], itemID, false)
}
