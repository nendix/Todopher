package funcs

import (
	"fmt"

	utils "github.com/nendix/Todopher/internal/utils"
)

// MarkTodos marks multiple todos as completed based on their IDs in the specified file.
func MarkTodos(filePath string, ids []uint8) error {
	// Read all Todos from the file using the helper function.
	todos, err := utils.ReadAllTodos(filePath)
	if err != nil {
		return fmt.Errorf("failed to read todos: %v", err)
	}

	// Create a set of IDs to mark for quick lookup.
	idSet := make(map[uint8]struct{})
	for _, id := range ids {
		idSet[id] = struct{}{}
	}

	// Flag to check if at least one Todo was marked.
	marked := false

	// Iterate through the todos and mark the specified ones as completed.
	for i, todo := range todos {
		if _, exists := idSet[todo.ID]; exists {
			if !todos[i].Completed {
				todos[i].Completed = true
				marked = true
			}
		}
	}

	if !marked {
		return fmt.Errorf("no matching todos found to mark")
	}

	// Write the updated Todos back to the file.
	err = utils.WriteAllTodos(filePath, todos)
	if err != nil {
		return fmt.Errorf("failed to write updated todos: %v", err)
	}

	return nil
}
