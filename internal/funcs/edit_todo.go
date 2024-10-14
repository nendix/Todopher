package funcs

import (
	"fmt"
	"strings"

	utils "github.com/nendix/Todopher/internal/utils"
)

// EditTodo modifies the description and/or due date of a todo with the specified ID.
func EditTodo(filePath string, id uint8, newTodo string, newDate string) error {
	// Read all Todos from the file.
	todos, err := utils.ReadAllTodos(filePath)
	if err != nil {
		return fmt.Errorf("failed to read todos: %v", err)
	}

	// Flag to check if the todo with the specified ID exists.
	found := false

	// Iterate through the todos to find and update the specified one.
	for i, todo := range todos {
		if todo.ID == id {
			found = true

			// Update the description if a new one is provided.
			if strings.TrimSpace(newTodo) != "" {
				todos[i].Description = newTodo
			}

			// Update the date if a new one is provided.
			if strings.TrimSpace(newDate) != "" {
				parsedDate, err := utils.ParseDate(newDate)
				if err != nil {
					return fmt.Errorf("failed to parse new date: %v", err)
				}
				todos[i].Date = parsedDate
			}

			break
		}
	}

	if !found {
		return fmt.Errorf("todo with ID %03d not found", id)
	}

	// Write the updated Todos back to the file.
	err = utils.WriteAllTodos(filePath, todos)
	if err != nil {
		return fmt.Errorf("failed to write updated todos: %v", err)
	}

	return nil
}
