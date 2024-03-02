package todo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"inodaf/todo/utils"
	"inodaf/todo/internal/config"
)

func Edit() {
	// @TODO: Fallback to last item in case the ID was not specified.
	if len(os.Args) <= 2 {
		fmt.Println("Edit: Please specify the item ID\nExample: `$ todo edit 12`")
		return
	}

	// Get the item ID from the CLI arguments.
	// $ todo edit <id>
	itemID, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Edit: Please use a valid number")
		return
	}

	// Access all the items.
	items := utils.GetItems(config.DatabasePath)
	if itemID > len(items) {
		fmt.Println("Edit: The item does not exists.")
		return
	}

	item := items[itemID]

	// Generate the text content to be put in the temporary file.
	content := fmt.Sprintf("%s\n%s", item.Title, item.Description)
	tempFileName := fmt.Sprintf("tmp-edit-%d.md", itemID)

	// Create a new temp file for editing.
	file, err := os.OpenFile(tempFileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("Edit: Error while creating the temporary editable file.")
		return
	}

	// When function is finished/errored, make sure to close and remove the file.
	defer os.Remove(tempFileName)
	defer file.Close()

	// Write contents to the file.
	if _, err := file.WriteString(content); err != nil {
		fmt.Printf("Edit: Error while writing contents to file. Err: %s\n", err)
		return
	}

	// Open the temp file using Vim.
	cmd := exec.Command("vim", tempFileName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Edit: Error while opening editor.")
		return
	}

	// Reopen the file to get the last contents.
	file, err = os.Open(tempFileName)
	if err != nil {
		fmt.Println("Edit: Error while opening the temporary file.")
		return
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	// The first line of the file is considered the item "title"
	// and it is "required" to create a new item.
	fileScanner.Scan()
	item.Title = fileScanner.Text()
	if len(item.Title) == 0 {
		fmt.Println("Error: Items must have a title.")
		return
	}

	// All the subsequent lines are considered the item "description"
	// and it is "not required" for the creation of the item.
	//
	// This also handles line-breaks by appending "\n"
	// when the line is blank or when reaching it's end.
	var descriptionBuilder strings.Builder
	for fileScanner.Scan() {
		var lineContent string = fileScanner.Text()

		if len(lineContent) == 0 {
			descriptionBuilder.WriteString("\n")
			continue
		}

		descriptionBuilder.WriteString(lineContent + "\n")
	}
	item.Description = descriptionBuilder.String()

	// Update the item in the store.
	items[itemID] = item

	// Convert the struct into a JSON string.
	data, err := json.Marshal(items)
	if err != nil {
		fmt.Println("Edit: Error while converting content to JSON.")
		return
	}

	utils.WriteItems(config.DatabasePath, data)
}
