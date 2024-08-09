package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	tgfuncs "github.com/nendix/TaskGopher/internal/tgfuncs"
	"github.com/nendix/TaskGopher/internal/utils"
)

func main() {
	// Ottieni la home directory dell'utente corrente
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	// Costruisci il percorso completo per il filename nella home directory
	filename := filepath.Join(homeDir, "todo", "todos.txt")

	// Crea il file se non esiste
	_, err = utils.CreateFileIfNotExists(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	if len(os.Args) < 2 {
		tgfuncs.Help()
		return
	}

	command := os.Args[1]

	switch command {
	case "help":
		tgfuncs.Help()

	case "add", "a":
		if len(os.Args) < 4 {
			fmt.Println("Usage: tg add [todo] [dd-mm-yy]")
			return
		}
		label := os.Args[2]
		due := os.Args[3]
		err := tgfuncs.AddToDo(filename, label, due)
		if err != nil {
			fmt.Println("Error adding todo:", err)
		} else {
			fmt.Println("Todo added successfully.")
		}

	case "edit", "e":
		if len(os.Args) < 5 {
			fmt.Println("Usage: tg edit [id] [new task] [new dd-mm-yy]")
			return
		}
		id, err := strconv.ParseUint(os.Args[2], 10, 8)
		if err != nil {
			fmt.Println("Invalid ID format:", err)
			return
		}
		newTask := os.Args[3]
		newDue := os.Args[4]
		err = tgfuncs.EditToDo(filename, uint8(id), newTask, newDue)
		if err != nil {
			fmt.Println("Error editing todo:", err)
		} else {
			fmt.Println("Todo edited successfully.")
		}

	case "mark", "m":
		if len(os.Args) < 3 {
			fmt.Println("Usage: tg mark [id1 id2 ...]")
			return
		}
		var ids []uint8
		for _, idStr := range os.Args[2:] {
			id, err := strconv.ParseUint(idStr, 10, 8)
			if err != nil {
				fmt.Println("Invalid ID format:", err)
				return
			}
			ids = append(ids, uint8(id))
		}
		err = tgfuncs.MarkToDos(filename, ids)
		if err != nil {
			fmt.Println("Error marking todo:", err)
		} else {
			fmt.Println("Todo marked as completed.")
		}

	case "unmark", "u":
		if len(os.Args) < 3 {
			fmt.Println("Usage: tg unmark [id1 id2 ...]")
			return
		}
		var ids []uint8
		for _, idStr := range os.Args[2:] {
			id, err := strconv.ParseUint(idStr, 10, 8)
			if err != nil {
				fmt.Println("Invalid ID format:", err)
				return
			}
			ids = append(ids, uint8(id))
		}
		err = tgfuncs.UnmarkToDos(filename, ids)
		if err != nil {
			fmt.Println("Error unmarking todo:", err)
		} else {
			fmt.Println("Todo unmarked.")
		}

	case "list", "ls":
		err := tgfuncs.ListToDos(filename)
		if err != nil {
			fmt.Println("Error listing todos:", err)
		}

	case "search", "s":
		if len(os.Args) < 3 {
			fmt.Println("Usage: tg search [keyword]")
			return
		}
		keyword := os.Args[2]
		err := tgfuncs.SearchToDos(filename, keyword)
		if err != nil {
			fmt.Println("Error searching todos:", err)
		}

	case "sort":
		if len(os.Args) < 3 {
			fmt.Println("Usage: tg sort [by_date|by_status]")
			return
		}
		criteria := os.Args[2]
		err := tgfuncs.SortToDos(filename, criteria)
		if err != nil {
			fmt.Println("Error sorting todos:", err)
		}

	case "delete", "d":
		if len(os.Args) < 3 {
			fmt.Println("Usage: tg delete [id1 id2 ...]")
			return
		}
		var ids []uint8
		for _, idStr := range os.Args[2:] {
			id, err := strconv.ParseUint(idStr, 10, 8)
			if err != nil {
				fmt.Println("Invalid ID format:", err)
				return
			}
			ids = append(ids, uint8(id))
		}
		err = tgfuncs.DeleteToDos(filename, ids)
		if err != nil {
			fmt.Println("Error deleting todos:", err)
		} else {
			fmt.Println("Todos deleted.")
		}

	default:
		fmt.Println("Unknown command:", command)
		tgfuncs.Help()
	}
}
