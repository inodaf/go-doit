package main

import (
	"fmt"
	"log"
	"os"

	"inodaf/todo/internal/pkg/cli"
	"inodaf/todo/internal/pkg/database"
)

func main() {
	db, err := database.NewSQLiteStore()
	if err != nil {
		log.Fatal("todo: could not connect to database")
	}
	defer db.Close()
	database.DB = db

	// Apply the Database Schema
	err = database.Prepare(db)
	if err != nil {
		log.Fatal(err.Error())
	}

	if len(os.Args) <= 1 {
		cli.HandleList()
		return
	}

	switch os.Args[1] {
	case "help":
		fmt.Println("Available actions: \n- add\n- list \n- edit\n- done\n- undone\n- view\n- remove")
		return
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
