package tgfuncs

import (
	"fmt"

	utils "github.com/nendix/TaskGopher/internal/utils"
)

// DeleteToDos elimina una lista di attività dal file specificato
func DeleteToDos(filePath string, ids []uint8) error {
	// Read all ToDos from the file.
	todos, err := utils.ReadAllToDos(filePath)
	if err != nil {
		return fmt.Errorf("failed to read todos: %v", err)
	}

	// Create a set of IDs to delete for quick lookup.
	idSet := make(map[uint8]struct{})
	for _, id := range ids {
		idSet[id] = struct{}{}
	}

	// Filter out the ToDos that need to be deleted.
	var updatedToDos []utils.ToDo
	for _, todo := range todos {
		if _, exists := idSet[todo.ID]; !exists {
			updatedToDos = append(updatedToDos, todo)
		}
	}

	// Reassign IDs to maintain continuity.
	for i := range updatedToDos {
		updatedToDos[i].ID = uint8(i + 1)
	}

	// Write the updated ToDos back to the file.
	err = utils.WriteAllToDos(filePath, updatedToDos)
	if err != nil {
		return fmt.Errorf("failed to write updated todos: %v", err)
	}

	return nil
}
