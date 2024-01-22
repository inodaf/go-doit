package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const DatabasePath = "./store/current.json"
const TempFileName = "tmp.md"

// todo edit 1
// This will open VIM
// Create temp edit file
// Open temp edit file in VIM
// When VIM is closed, save temp edit file
//
// Read temp edit file
// Delete temp edit file
//
// Session-based to dos with authentication
// Persistent session and "REPL" session
// Encrypted DB
//
// Separate Clients (CLI, Web, etc...)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please, specify the action: add, list, update")
		return
	}

	switch os.Args[1] {
	case "list":
		list()
		return
	case "view":
		view()
		return
	case "done":
		markAsDone()
		return
	case "undone":
		markAsUndone()
		return
	case "add":
	default:
	}

	// Start: Spawns the vim process and save the tmp.md file.
	cmd := exec.Command("vim", TempFileName)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	// End

	// Start: Read the tmp.txt file and get it's content
	file, err := os.Open(TempFileName)
	if err != nil {
		panic(err)
	}

	defer os.Remove(TempFileName)
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

	add(title, description)
	list()
}
