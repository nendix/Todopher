package tgfuncs

import (
	"fmt"
	"sort"
	"strings"
	"time"

	utils "github.com/nendix/TaskGopher/internal/utils"
)

// SortToDos sorts the todos in the specified file based on the given criteria ("date" or "status") and reassigns IDs.
func SortToDos(filePath, criteria string) error {
	// Read all ToDos from the file using the helper function.
	todos, err := utils.ReadAllToDos(filePath)
	if err != nil {
		return fmt.Errorf("failed to read todos: %v", err)
	}

	// Define sorting based on criteria.
	switch strings.ToLower(criteria) {
	case "date", "d":
		sort.Slice(todos, func(i, j int) bool {
			// Convert Date struct to time.Time for comparison.
			dateI := time.Date(todos[i].Date.Year, time.Month(todos[i].Date.Month), todos[i].Date.Day, 0, 0, 0, 0, time.UTC)
			dateJ := time.Date(todos[j].Date.Year, time.Month(todos[j].Date.Month), todos[j].Date.Day, 0, 0, 0, 0, time.UTC)
			return dateI.Before(dateJ)
		})
	case "status", "s":
		sort.Slice(todos, func(i, j int) bool {
			// Incomplete todos come before completed ones.
			return !todos[i].Completed && todos[j].Completed
		})
	default:
		return fmt.Errorf("unknown sorting criteria: %s", criteria)
	}

	// Reassign IDs to maintain continuity.
	for idx := range todos {
		todos[idx].ID = uint8(idx + 1)
	}

	// Write the sorted ToDos back to the file.
	err = utils.WriteAllToDos(filePath, todos)
	if err != nil {
		return fmt.Errorf("failed to write sorted todos: %v", err)
	}

	return nil
}
