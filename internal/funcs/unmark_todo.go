package funcs

import (
	"fmt"

	utils "github.com/nendix/Todopher/internal/utils"
)

// UnmarkTodos deselects multiple todos in the specified file based on their IDs.
func UnmarkTodos(filePath string, ids []uint8) error {
	// Read all Todos from the file using the helper function.
	todos, err := utils.ReadAllTodos(filePath)
	if err != nil {
		return fmt.Errorf("failed to read todos: %v", err)
	}

	// Create a set of IDs to unmark for quick lookup.
	idSet := make(map[uint8]struct{})
	for _, id := range ids {
		idSet[id] = struct{}{}
	}

	// Flag to check if at least one Todo was unmarked.
	unmarked := false

	// Iterate through the todos and unmark the specified ones.
	for i, todo := range todos {
		if _, exists := idSet[todo.ID]; exists {
			if todos[i].Completed {
				todos[i].Completed = false
				unmarked = true
			}
		}
	}

	if !unmarked {
		return fmt.Errorf("no matching todos found to unmark")
	}

	// Write the updated Todos back to the file.
	err = utils.WriteAllTodos(filePath, todos)
	if err != nil {
		return fmt.Errorf("failed to write updated todos: %v", err)
	}

	return nil
}
