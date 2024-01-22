package main

import (
	"flag"
	"os"

	"github.com/inodaf/todo/utils"
)

func list() {

	// When no arguments - flags - are provided, assume
	// to list only "pending" items.
	if len(os.Args) <= 2 {
		listPendingItems()
		return
	}

	// Custom "flag" set, as the default "flag.Parse" parses
	// the second argument from "os.Args". Given our CLI interface,
	// we need to parse the third argument instead.
	//
	// $ todo list -a
	var options *flag.FlagSet = flag.NewFlagSet(os.Args[2], flag.ExitOnError)
	var listAll *bool = options.Bool("a", false, "List pending and done items")
	var onlyDone *bool = options.Bool("c", false, "List done items")
	options.Parse(os.Args[2:])

	if *onlyDone {
		listDoneItems()
		return
	}

	listPendingItems()

	if !*listAll && !*onlyDone {
		return
	}

	listDoneItems()
}

func listDoneItems() {
	for index, item := range utils.GetItems(DatabasePath) {
		if item.DoneAt != "" {
			utils.PrintItem(&item, index, false)
		}
	}
}

func listPendingItems() {
	for index, item := range utils.GetItems(DatabasePath) {
		if item.DoneAt != "" {
			continue
		}
		utils.PrintItem(&item, index, false)
	}
}
