package todo

import (
	"fmt"
	"os"
	"strconv"

	"inodaf/todo/internal/config"
	"inodaf/todo/utils"
)

func View() {
	if len(os.Args) <= 2 {
		fmt.Println("View: Please specify the item ID\nExample: `$ todo view 12`")
		return
	}

	itemID, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("View: Please use a valid number")
		return
	}

	items := utils.GetItems(config.DatabasePath)

	if itemID > len(items) {
		fmt.Println("View: Not valid")
		return
	}

	// Solve ID mapping with a proper DB instead of a JSON.
	item := items[itemID]
	utils.PrintItem(&item, itemID, true)
}
