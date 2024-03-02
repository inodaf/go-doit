package todo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"inodaf/todo/internal/config"
	"inodaf/todo/internal/models"
	"inodaf/todo/utils"
)

func Add() {
	// Start: Spawns the vim process and save the tmp.md file.
	cmd := exec.Command("vim", config.TempFileName)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Print("Add: \"Editor\" process did not finish successfully")
		return
	}
	// End

	// Start: Read the tmp.txt file and get it's content
	file, err := os.Open(config.TempFileName)
	if err != nil {
		fmt.Print("Add: Error while opening the temporary file.")
		return
	}

	defer os.Remove(config.TempFileName)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Scan()

	// The first line of the file is considered the item "title"
	// and it is "required" to create a new item.
	var title string = fileScanner.Text()
	if len(title) == 0 {
		fmt.Print("Error: Items must have a title.\n")
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

	item := *models.NewItem()
	item.Title = title
	item.Description = description

	items := utils.GetItems(config.DatabasePath)
	items = append(items, item)

	data, err := json.Marshal(items)
	if err != nil {
		panic(err)
	}

	utils.WriteItems(config.DatabasePath, data)
}
