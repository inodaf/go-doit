package cli

import (
	"fmt"
	"inodaf/todo/internal/config"
	"inodaf/todo/internal/pkg/models"
	"strings"
	"time"
)

func RenderTemplate(item *models.Item, withDetails bool) string {
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
	if len(item.DoneAt) != 0 {
		completed = "x"
		doneAt, _ := time.Parse(time.DateTime, item.DoneAt)
		template.WriteString(fmt.Sprintf("Done at %s\n", doneAt.Format(config.DisplayTimeFormat)))
	}

	// Append the line "Created at <time>"
	// for all rendering cases.
	createdAt, _ := time.Parse(time.DateTime, item.CreatedAt)
	template.WriteString(fmt.Sprintf("Created at %s\n", createdAt.Format(config.DisplayTimeFormat)))

	// When rendering the view with details,
	// show the item's description.
	//
	// E.g.: [x] #1: My item title
	// Created at 22 Jan 24 18:02 CET
	// Updated at 24 Jan 24 11:37 CET
	//
	// This is the item description.
	if withDetails {
		if len(item.UpdatedAt) != 0 {
			updatedAt, _ := time.Parse(time.DateTime, item.UpdatedAt)
			template.WriteString(fmt.Sprintf("Updated at %s\n", updatedAt.Format(config.DisplayTimeFormat)))
		}
		template.WriteString(fmt.Sprintf("\n%s\n", item.Description))
	}

	return fmt.Sprintf(template.String(), completed, item.Id, item.Title)
}

func PrintItem(item *models.Item, compact bool) {
	fmt.Print(RenderTemplate(item, compact))
}
