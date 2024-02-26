package main

import (
	"fmt"
	"os"
)

const DatabasePath = "./store/current.json"
const TempFileName = "tmp.md"

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please, use one of the available actions: \n- add\n- list \n- edit\n- done\n- undone\n- view")
		return
	}

	switch os.Args[1] {
	case "list":
		list()
		return
	case "view":
		view()
		return
	case "done":
		markAsDone()
		return
	case "undone":
		markAsUndone()
		return
	case "edit":
		edit()
		return
	case "add":
		add()
		list()
		return
	default:
		list()
		return
	}
}
