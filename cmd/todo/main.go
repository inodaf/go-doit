package main

import (
	"fmt"
	"os"

	"inodaf/todo/internal/pkg/cli"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please, use one of the available actions: \n- add\n- list \n- edit\n- done\n- undone\n- view")
		return
	}

	switch os.Args[1] {
	case "list":
		cli.HandleList()
		return
	case "view":
		cli.HandleView()
		return
	case "done":
		cli.HandleMarkDone()
		return
	case "undone":
		cli.HandleMarkUndone()
		return
	case "edit":
		cli.HandleEdit()
		return
	case "add":
		cli.HandleAdd()
		cli.HandleList()
		return
	case "remove":
		cli.HandleRemove()
		return
	default:
		cli.HandleList()
		return
	}
}
