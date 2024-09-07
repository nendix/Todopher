package tg

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nendix/TaskGopher/internal/utils"
)

const (
	todoDirName     = "todo"
	envFileName     = ".env"
	initialFileName = "todos.txt"
)

// Init crea la cartella ~/todo e il file ~/todo/.env con la variabile TODO_FILE=todos.txt
func Init() error {
	// Costruisci il percorso della cartella todo e del file .env
	todoDirPath, err := utils.GetTodoDir()
	if err != nil {
		return fmt.Errorf("error getting todo directory: %w", err)
	}
	envFilePath := filepath.Join(todoDirPath, envFileName)

	// Controlla se la cartella todo esiste, altrimenti la crea
	info, err := os.Stat(todoDirPath)
	if err == nil {
		if !info.IsDir() {
			return fmt.Errorf("file exists with the name '%s', but it's not a directory", todoDirPath)
		}
		fmt.Println("Directory already exists.")
	} else if os.IsNotExist(err) {
		err = os.MkdirAll(todoDirPath, 0755)
		if err != nil {
			return fmt.Errorf("error creating todo directory: %w", err)
		}
		fmt.Println("Directory created successfully.")
	} else {
		return fmt.Errorf("error checking todo directory: %w", err)
	}

	// Crea il file .env se non esiste
	_, err = os.Stat(envFilePath)
	if os.IsNotExist(err) {
		file, err := os.Create(envFilePath)
		if err != nil {
			return fmt.Errorf("error creating .env file: %w", err)
		}
		defer file.Close()

		// Scrivi la variabile d'ambiente TODO_FILE=todos.txt
		_, err = file.WriteString("TODO_FILE=" + initialFileName + "\n")
		if err != nil {
			return fmt.Errorf("error writing to .env file: %w", err)
		}
		fmt.Println("File .env created and initialized successfully.")
	} else if err != nil {
		return fmt.Errorf("error checking .env file: %w", err)
	}

	return nil
}
