package cli

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"inodaf/todo/internal/config"
	"inodaf/todo/internal/todos"
)

func HandleView() {
	if len(os.Args) <= 2 {
		fmt.Println("View: Please specify the item ID\nExample: `$ todo view 12`")
		return
	}

	itemID, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("View: Please use a valid number")
		return
	}

	item, err := todos.View(itemID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	PrintItem(item, true)
}

func HandleList() {
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

	err := options.Parse(os.Args[2:])
	if err != nil {
		fmt.Println("List: Failed to parse flags.")
		return
	}

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

func HandleAdd() {
	tempDir, err := os.MkdirTemp("", "godoit")
	if err != nil {
		log.Println("Add: Could not create temporary file.", err.Error())
		return
	}
	defer os.RemoveAll(tempDir)

	tmpFilePath := filepath.Join(tempDir, config.TempFileName)

	// Spawns the $EDITOR process and save the temp file.
	cmd := exec.Command("vim", tmpFilePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Println("Add: \"Editor\" process did not finish successfully")
		return
	}

	// Read the temp file and get its content
	file, err := os.Open(tmpFilePath)
	if err != nil {
		log.Println("Add: Error while opening the temporary file.")
		return
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Scan()

	// The first line of the file is considered the item "title"
	// and it is "required" to create a new item.
	var title string = fileScanner.Text()
	if len(title) == 0 {
		log.Println("Error: Items must have a title.")
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
	var description string = descriptionBuilder.String()

	err = todos.Add(todos.AddInput{Title: title, Description: description})
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func HandleEdit() {
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

	// Access the specified item given its ID.
	item, err := todos.View(itemID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Generate the text content to be put in the temporary file.
	tempFileName := fmt.Sprintf("tmp-edit-%d.md", itemID)
	content := fmt.Sprintf("%s\n%s", item.Title, item.Description)

	// Create a new temp file for editing.
	tempDir, err := os.MkdirTemp("", "godoit")
	if err != nil {
		log.Println("Add: Could not create the temporary directory.")
		return
	}
	defer os.RemoveAll(tempDir)

	filePath := filepath.Join(tempDir, tempFileName)
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("Edit: Error while creating the temporary editable file.")
		return
	}

	// When function is finished/errored, make sure to close and remove the file.
	defer file.Close()

	// Write contents to the file.
	if _, err := file.WriteString(content); err != nil {
		fmt.Printf("Edit: Error while writing contents to file. Err: %s\n", err)
		return
	}

	// Open the temp file using Vim.
	cmd := exec.Command("vim", filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Edit: Error while opening editor.")
		return
	}

	// Reopen the file to get the last contents.
	file, err = os.Open(filePath)
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

	err = todos.Edit(todos.EditInput{Item: item})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("edit: changes saved")
	PrintItem(item, false)
}

func HandleMarkDone() {
	if len(os.Args) <= 2 {
		fmt.Println("Mark as Done: Please specify the item ID\nExample: `$ todo done 12`.")
		return
	}

	itemID, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Mark as Done: Please specify a valid item ID.")
		return
	}

	item, err := todos.MarkAsDone(itemID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	PrintItem(item, false)
}

func HandleMarkUndone() {
	if len(os.Args) <= 2 {
		fmt.Println("Mark as Undone: Please specify the item ID\nExample: `$ todo undone 12`.")
		return
	}

	itemID, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Mark as Undone: Please specify a valid item ID.")
		return
	}

	item, err := todos.MarkAsUndone(itemID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	PrintItem(item, false)
}

func HandleRemove() {
	if len(os.Args) <= 2 {
		fmt.Println("Remove: Please specify the item IDs.")
		return
	}

	itemID, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Remove: Please specify a valid item ID.")
		return
	}

	err = todos.Remove(itemID, false)
	if err == todos.ErrItemIsNotDone {
		var confirmation string

		fmt.Println("Item is not done yet. Confirm deletion?: y/n")
		fmt.Scan(&confirmation)

		if strings.ToLower(confirmation) != "y" {
			fmt.Println("Ok, not removing!")
			return
		}

		err = todos.Remove(itemID, true)
	}

	if err != nil {
		fmt.Printf("Remove: Failed to remove item. Cause: %s", err.Error())
		return
	}

	fmt.Printf("Removed item %d. \n", itemID)
}

// Helpers
func listDoneItems() {
	items, err := todos.ListDoneItems()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(items) == 0 {
		fmt.Println("📋 Completed tasks will appear here")
		return
	}

	for _, item := range items {
		PrintItem(item, false)
	}
}

func listPendingItems() {
	items, err := todos.ListPendingItems()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(items) == 0 {
		fmt.Println("✅ No tasks")
		return
	}

	for _, item := range items {
		PrintItem(item, false)
	}
}
