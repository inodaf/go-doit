package main

import (
	"fmt"
	"os"

	"inodaf/todo/internal/handlers/cli"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please, use one of the available actions: \n- add\n- list \n- edit\n- done\n- undone\n- view")
		return
	}

	switch os.Args[1] {
	case "list":
		cli.List()
		return
	case "view":
		cli.View()
		return
	case "done":
		cli.MarkDone()
		return
	case "undone":
		cli.MarkUndone()
		return
	case "edit":
		cli.Edit()
		return
	case "add":
		cli.Add()
		cli.List()
		return
	case "remove":
		cli.Remove()
		return
	default:
		cli.List()
		return
	}
}
