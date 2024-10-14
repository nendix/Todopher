package funcs

import (
	"errors"
	"fmt"
	"strings"

	utils "github.com/nendix/Todopher/internal/utils"
)

// SearchTodos searches for todos containing the specified keyword in the given file.
func SearchTodos(filePath, keyword string) error {
	// Read all Todos from the file using the helper function.
	todos, err := utils.ReadAllTodos(filePath)
	if err != nil {
		return fmt.Errorf("failed to read todos: %v", err)
	}

	// Flag to check if any Todo matches the keyword.
	found := false

	// Iterate through each Todo and check if the Description contains the keyword.
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
