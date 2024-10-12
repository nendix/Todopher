package tgfuncs

import (
	"fmt"
	"path/filepath"
	"strings"

	utils "github.com/nendix/TaskGopher/internal/utils"
)

// ListToDos prints the list of all todos from the specified file.
func ListToDos(filePath string) error {
	// Read the current list name using the utility function.
	listName, err := utils.ReadCurrentList()
	if err != nil {
		return fmt.Errorf("error reading current list: %v", err)
	}

	// Remove the .txt extension from the list name for display purposes.
	listName = strings.TrimSuffix(listName, filepath.Ext(listName))

	if listName != "" {
		fmt.Println("List:", listName)
	}

	todos, err := utils.ReadAllToDos(filePath)
	if err != nil {
		return fmt.Errorf("error reading todos: %v", err)
	}

	// Check if there are no todos to display.
	if len(todos) == 0 {
		fmt.Println("No todos found.")
		return nil
	}

	// Iterate through each ToDo and print it in a formatted manner.
	for _, todo := range todos {
		status := "[ ]"
		if todo.Completed {
			status = "[✓]"
		}
		fmt.Printf("%03d %s %s - %s\n", todo.ID, status, todo.Description, todo.Date.String())
	}

	return nil
}
