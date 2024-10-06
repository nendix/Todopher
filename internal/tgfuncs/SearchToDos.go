package tgfuncs

import (
	"errors"
	"fmt"
	"strings"

	utils "github.com/nendix/TaskGopher/internal/utils"
)

// SearchToDos searches for todos containing the specified keyword in the given file.
func SearchToDos(filePath, keyword string) error {
	// Read all ToDos from the file using the helper function.
	todos, err := utils.ReadAllToDos(filePath)
	if err != nil {
		return fmt.Errorf("failed to read todos: %v", err)
	}

	// Flag to check if any ToDo matches the keyword.
	found := false

	// Iterate through each ToDo and check if the Description contains the keyword.
	for _, todo := range todos {
		if strings.Contains(strings.ToLower(todo.Description), strings.ToLower(keyword)) {
			fmt.Println(todo.String())
			found = true
		}
	}

	if !found {
		return errors.New("no todos found with the given keyword")
	}

	return nil
}
