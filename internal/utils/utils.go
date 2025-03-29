package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
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
	return filepath.Join(homeDir, "Todos"), nil
}

func GetEnvFilePath() (string, error) {
	todoDir, err := GetTodoDir()
	if err != nil {
		return "", err
	}

	configFilePath := filepath.Join(todoDir, ".env")
	// Check if .env exists
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		return "", fmt.Errorf(".env does not exist")
	}

	return configFilePath, nil
}

// ReadCurrentList legge il file .env e restituisce il valore di TODO_FILE
func ReadCurrentList() (string, error) {
	todoDir, err := GetTodoDir()
	if err != nil {
		return "", err
	}

	// Carica il file .env
	envFilePath := filepath.Join(todoDir, ".env")
	err = godotenv.Load(envFilePath)
	if err != nil {
		return "", fmt.Errorf("error loading .env file: %w", err)
	}

	// Ottieni la variabile d'ambiente TODO_FILE
	fileName := os.Getenv("TODO_FILE")
	if fileName == "" {
		return "", fmt.Errorf("TODO_FILE not set in .env file")
	}

	return fileName, nil
}

func GetTodoFilePath() (string, error) {
	todoDir, err := GetTodoDir()
	if err != nil {
		return "", fmt.Errorf("error getting todo directory: %w", err) // TODO: err
	}
	fileName, err := ReadCurrentList()
	if err != nil {
		return "", fmt.Errorf("error getting current list: %w", err) // TODO: err
	}
	return filepath.Join(todoDir, fileName), nil
}
