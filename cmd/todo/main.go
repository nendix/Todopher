package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

	// Crea il file se non esiste e controlla se è stato creato
	created, err := utils.CreateFileIfNotExists(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	if created {
		fmt.Println("File created successfully in home directory.")
	}

	for {
		utils.ShowMenu()
		choice := utils.ReadChoice()

		switch choice {
		case 1:
			// Aggiungi un nuovo todo
			fmt.Print("Enter todo label: ")
			label := utils.ReadString()
			fmt.Print("Enter due date (dd/mm/yy): ")
			due := utils.ReadString()

			err := crud.AddToDo(filename, label, due)
			if err != nil {
				fmt.Println("Error adding todo:", err)
			}

		case 2:
			fmt.Print("Enter todo ID to mark as completed: ")
			idStr := utils.ReadString()
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				fmt.Println("Invalid ID. Please enter a valid number.")
				break
			}

			err = crud.MarkToDo(filename, uint8(id))
			if err != nil {
				fmt.Println("Error marking todo as completed:", err)
			} else {
				fmt.Println("Todo marked as completed.")
			}
		case 3:
			// todos = unmarkTodo() // Aggiorna la variabile todos
		case 4:
			// todos = deleteTodo() // Aggiorna la variabile todos
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
