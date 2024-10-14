package funcs

import (
	"fmt"

	utils "github.com/nendix/Todopher/internal/utils"
)

// DeleteTodos elimina una lista di attività dal file specificato
func DeleteTodos(filePath string, ids []uint8) error {
	// Read all Todos from the file.
	todos, err := utils.ReadAllTodos(filePath)
	if err != nil {
		return fmt.Errorf("failed to read todos: %v", err)
	}

	// Create a set of IDs to delete for quick lookup.
	idSet := make(map[uint8]struct{})
	for _, id := range ids {
		idSet[id] = struct{}{}
	}

	// Filter out the Todos that need to be deleted.
	var updatedTodos []utils.Todo
	for _, todo := range todos {
		if _, exists := idSet[todo.ID]; !exists {
			updatedTodos = append(updatedTodos, todo)
		}
	}

	// Reassign IDs to maintain continuity.
	for i := range updatedTodos {
		updatedTodos[i].ID = uint8(i + 1)
	}

	// Write the updated Todos back to the file.
	err = utils.WriteAllTodos(filePath, updatedTodos)
	if err != nil {
		return fmt.Errorf("failed to write updated todos: %v", err)
	}

	return nil
}
