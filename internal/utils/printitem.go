package utils

import (
	"fmt"
	"strings"

	"inodaf/todo/internal/models"
)

func RenderTemplate(item *models.Item, index int, withDetails bool) string {
	var template strings.Builder
	var completed string

	// Adds a separator between items and
	// define the template for the item first line.
	//
	// E.g.: [x] #1: My item title
	template.WriteString("\n----\n")
	template.WriteString("[%s] #%d: %s\n")

	// When item is marked as "done",
	// add the "x" to symbolize completion and
	// append the line "Done at <time>".
	if item.DoneAt != "" {
		completed = "x"
		template.WriteString(fmt.Sprintf("Done at %s\n", item.DoneAt))
	}

	// Append the line "Created at <time>"
	// for all rendering cases.
	template.WriteString(fmt.Sprintf("Created at %s\n", item.CreatedAt))

	// When rendering the view with details,
	// show the item's description.
	//
	// E.g.: [x] #1: My item title
	// Created at 22 Jan 24 18:02 CET
	//
	// This is the item description.
	if withDetails {
		if len(item.UpdatedAt) != 0 {
			template.WriteString(fmt.Sprintf("Updated at %s\n", item.UpdatedAt))
		}
		template.WriteString(fmt.Sprintf("\n%s\n", item.Description))
	}

	return fmt.Sprintf(template.String(), completed, index, item.Title)
}

func PrintItem(item *models.Item, index int, compact bool) {
	fmt.Print(RenderTemplate(item, index, compact))
}
