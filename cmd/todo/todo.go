package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"todo/internal/utils"
	"todo/pkg/crud"
)

func main() {
	// Ottieni la home directory dell'utente corrente
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	// Costruisci il percorso completo per il file nella home directory
	filename := filepath.Join(homeDir, "todo", "todos.txt")

	// Crea il file se non esiste
	_, err = utils.CreateFileIfNotExists(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	if len(os.Args) < 2 {
		utils.PrintHelp()
		return
	}

	command := os.Args[1]

	switch command {
	case "add", "a":
		if len(os.Args) < 4 {
			fmt.Println("Usage: todo add [todo] [dd-mm-yy]")
			return
		}
		label := os.Args[2]
		due := os.Args[3]
		err := crud.AddToDo(filename, label, due)
		if err != nil {
			fmt.Println("Error adding todo:", err)
		} else {
			fmt.Println("Todo added successfully.")
		}

	case "mark", "m":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo mark [id]")
			return
		}
		id, err := strconv.ParseUint(os.Args[2], 10, 8)
		if err != nil {
			fmt.Println("Invalid ID format:", err)
			return
		}
		err = crud.MarkToDo(filename, uint8(id))
		if err != nil {
			fmt.Println("Error marking todo:", err)
		} else {
			fmt.Println("Todo marked as completed.")
		}

	case "unmark", "u":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo unmark [id]")
			return
		}
		id, err := strconv.ParseUint(os.Args[2], 10, 8)
		if err != nil {
			fmt.Println("Invalid ID format:", err)
			return
		}
		err = crud.UnmarkToDo(filename, uint8(id))
		if err != nil {
			fmt.Println("Error unmarking todo:", err)
		} else {
			fmt.Println("Todo unmarked.")
		}

	case "list", "ls":
		err := crud.ListToDos(filename)
		if err != nil {
			fmt.Println("Error listing todos:", err)
		}

	case "delete", "d":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo delete [id]")
			return
		}
		id, err := strconv.ParseUint(os.Args[2], 10, 8)
		if err != nil {
			fmt.Println("Invalid ID format:", err)
			return
		}
		err = crud.DeleteToDo(filename, uint8(id))
		if err != nil {
			fmt.Println("Error deleting todo:", err)
		} else {
			fmt.Println("Todo deleted.")
		}

	default:
		fmt.Println("Unknown command:", command)
		utils.PrintHelp()
	}
}
