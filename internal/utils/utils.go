package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFileIfNotExists(filename string) (bool, error) {
	// Crea il file se non esiste
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return false, fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	// Controlla se il file è stato effettivamente creato
	fi, err := file.Stat()
	if err != nil {
		return false, err
	}

	// Controlla se il file era già presente
	return fi.Size() == 0, nil
}

// Function to get the path of the todo directory
func GetTodoDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, "todo"), nil
}

func GetConfigFilePath() (string, error) {
	todoDir, err := GetTodoDir()
	if err != nil {
		return "", err
	}

	configFilePath := filepath.Join(todoDir, "CURRENT_TODO_FILE.txt")
	// Check if CURRENT_TODO_FILE.txt exists
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		return "", fmt.Errorf("CURRENT_TODO_FILE.txt does not exist")
	}

	return configFilePath, nil
}

// Function to ensure CURRENT_TODO_FILE.txt exists and contains todos.txt
func ensureCurrentFile() error {
	configFilePath, err := GetConfigFilePath()
	if err != nil {
		return err
	}

	// Check if CURRENT_TODO_FILE.txt exists
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		// If it doesn't exist, create it and write 'todos.txt'
		err := os.WriteFile(configFilePath, []byte("todos.txt"), 0644)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %v", configFilePath, err)
		}
	}

	return nil
}

// Function to read the current todo file from CURRENT_TODO_FILE.txt
func ReadCurrentList() (string, error) {
	configFilePath, err := GetConfigFilePath()
	if err != nil {
		return "", err
	}

	// Read the content of CURRENT_TODO_FILE.txt
	content, err := os.ReadFile(configFilePath)
	if err != nil {
		return "", err
	}

	// Return the filename as a string (removing any trailing newline characters)
	return string(content), nil
}
