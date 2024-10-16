package cli

import (
	"fmt"
	"os"
	"strconv"

	"github.com/nendix/Todopher/internal/funcs"
	"github.com/nendix/Todopher/internal/utils"
)

func HandleCLI() {
	command := os.Args[1]
	switch command {
	case "help":
		funcs.Help()
		return

	case "init":
		// Esegui il comando init per creare la cartella ~/todo e il file .env
		err := funcs.Init()
		if err != nil {
			fmt.Println("Error initializing:", err)
		} else {
			fmt.Println("Initialization completed successfully.")
		}
		return

	default:

		filePath, err := utils.GetTodoFilePath()
		if err != nil {
			fmt.Println(err)
			return
		}

		switch command {

		case "setlist", "sl":
			if len(os.Args) < 3 {
				fmt.Println("Usage: tg setlist [listname]")
				return
			}
			newFile := os.Args[2]

			err := funcs.SetList(newFile)
			if err != nil {
				fmt.Println("Error setting the list:", err)
			} else {
				fmt.Println("List set to:", newFile)
			}

		case "add", "a":
			if len(os.Args) < 4 {
				fmt.Println("Usage: tg add [todo] [dd-mm-yy]")
				return
			}
			label := os.Args[2]
			due := os.Args[3]
			err := funcs.AddTodo(filePath, label, due)
			if err != nil {
				fmt.Println("Error adding todo:", err)
			} else {
				fmt.Println("Todo added successfully.")
			}

		case "edit", "e":
			if len(os.Args) < 5 {
				fmt.Println("Usage: tg edit [id] [new todo] [new dd-mm-yy]")
				return
			}
			id, err := strconv.ParseUint(os.Args[2], 10, 8)
			if err != nil {
				fmt.Println("Invalid ID format:", err)
				return
			}
			newTodo := os.Args[3]
			newDue := os.Args[4]
			err = funcs.EditTodo(filePath, uint8(id), newTodo, newDue)
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
			err = funcs.MarkTodos(filePath, ids)
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
			err = funcs.UnmarkTodos(filePath, ids)
			if err != nil {
				fmt.Println("Error unmarking todo:", err)
			} else {
				fmt.Println("Todo unmarked.")
			}

		case "list", "ls":
			err := funcs.ListTodos(filePath)
			if err != nil {
				fmt.Println("Error listing todos:", err)
			}

		case "search", "s":
			if len(os.Args) < 3 {
				fmt.Println("Usage: tg search [keyword]")
				return
			}
			keyword := os.Args[2]
			err := funcs.SearchTodos(filePath, keyword)
			if err != nil {
				fmt.Println("Error searching todos:", err)
			}

		case "sort":
			if len(os.Args) < 3 {
				fmt.Println("Usage: tg sort [date|status]")
				return
			}
			criteria := os.Args[2]
			err := funcs.SortTodos(filePath, criteria)
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
			err = funcs.DeleteTodos(filePath, ids)
			if err != nil {
				fmt.Println("Error deleting todos:", err)
			} else {
				fmt.Println("Todos deleted.")
			}

		default:
			fmt.Println("Unknown command:", command)
			funcs.Help()
		}
	}
}
