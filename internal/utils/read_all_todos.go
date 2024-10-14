package utils

import (
	"bufio"
	"fmt"
	"os"
)

// ReadAllTodos reads all Todos from the specified file.
func ReadAllTodos(filePath string) ([]Todo, error) {
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []Todo{}, nil // If the file does not exist, return an empty list
		}
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var todos []Todo
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		todo, err := ParseTodo(line)
		if err != nil {
			// Log the error and skip the malformed line
			fmt.Printf("Warning: Skipping malformed line %d: %v\n", lineNumber, err)
			continue
		}
		todos = append(todos, todo)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return todos, nil
}
