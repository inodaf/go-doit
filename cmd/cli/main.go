package main

import (
	"fmt"
	"os"

	"inodaf/todo/internal/todo"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please, use one of the available actions: \n- add\n- list \n- edit\n- done\n- undone\n- view")
		return
	}

	switch os.Args[1] {
	case "list":
		todo.List()
		return
	case "view":
		todo.View()
		return
	case "done":
		todo.MarkAsDone()
		return
	case "undone":
		todo.MarkAsUndone()
		return
	case "edit":
		todo.Edit()
		return
	case "add":
		todo.Add()
		todo.List()
		return
	default:
		todo.List()
		return
	}
}
