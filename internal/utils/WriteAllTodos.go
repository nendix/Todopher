package utils

import (
	"bufio"
	"fmt"
	"os"
)

// WriteAllToDos scrive tutte le ToDos nel file specificato
func WriteAllToDos(filePath string, todos []ToDo) error {
	file, err := os.OpenFile(filePath, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("failed to open file for writing: %v", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, todo := range todos {
		_, err := writer.WriteString(todo.String() + "\n")
		if err != nil {
			return fmt.Errorf("failed to write todo ID %03d: %v", todo.ID, err)
		}
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to flush data to file: %v", err)
	}

	return nil
}
