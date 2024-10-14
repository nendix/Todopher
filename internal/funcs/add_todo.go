package funcs

import (
	"fmt"
	"os"

	utils "github.com/nendix/Todopher/internal/utils"
)

// AddTodo aggiunge un nuovo todo al file specificato
func AddTodo(filePath, todoDesc, dateStr string) error {
	// Parse the date.
	date, err := utils.ParseDate(dateStr)
	if err != nil {
		return fmt.Errorf("failed to parse date: %v", err)
	}

	// Get the last ID.
	lastID, err := readLastID(filePath)
	if err != nil {
		return fmt.Errorf("failed to read last ID: %v", err)
	}

	newID := lastID + 1

	// Create the new Todo.
	newTodo := utils.Todo{
		ID:          newID,
		Completed:   false, // Default to not completed.
		Description: todoDesc,
		Date:        date,
	}

	// Open the file in append mode, create if not exists.
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("failed to open file for appending: %v", err)
	}
	defer file.Close()

	// Write the new todo.
	_, err = file.WriteString(newTodo.String() + "\n")
	if err != nil {
		return fmt.Errorf("failed to write new todo: %v", err)
	}

	return nil
}

// readLastID reads the last ID used from the file.
func readLastID(filePath string) (uint8, error) {
	todos, err := utils.ReadAllTodos(filePath)
	if err != nil {
		return 0, fmt.Errorf("failed to read todos: %v", err)
	}

	if len(todos) == 0 {
		return 0, nil
	}

	return todos[len(todos)-1].ID, nil
}
