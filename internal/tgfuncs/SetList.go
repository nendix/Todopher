package tg

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nendix/TaskGopher/internal/utils"
)

// Function to get the path of the configuration file

// Function to save the current file in the configuration
func saveCurrentFile(listName string) error {
	configFile, err := utils.GetConfigFilePath()
	if err != nil {
		return err
	}

	// Write the filename to the config file
	err = os.WriteFile(configFile, []byte(listName), 0644)
	if err != nil {
		return err
	}

	return nil
}

// Function to check if the file exists in ~/todo/ or create it
func SetList(listName string) error {
	// Ottieni la directory ~/todo
	todoDir, err := utils.GetTodoDir()
	if err != nil {
		return fmt.Errorf("error getting todo directory: %v", err)
	}

	// Verifica se la directory ~/todo esiste e creala se non esiste
	if _, err := os.Stat(todoDir); os.IsNotExist(err) {
		err = os.MkdirAll(todoDir, 0755) // Creiamo la directory se non esiste
		if err != nil {
			return fmt.Errorf("failed to create todo directory: %v", err)
		}
		fmt.Println("Todo directory created successfully.")
	}

	// Aggiorna il file di configurazione con il nome del nuovo file
	err = saveCurrentFile(listName + ".txt")
	if err != nil {
		return fmt.Errorf("failed to update config: %v", err)
	}

	// Verifica se il file della lista esiste, se no, crealo
	listFilePath := filepath.Join(todoDir, listName+".txt")
	if _, err := os.Stat(listFilePath); os.IsNotExist(err) {
		// Creiamo il file della lista todo se non esiste
		file, err := os.Create(listFilePath)
		if err != nil {
			return fmt.Errorf("failed to create todo list file: %v", err)
		}
		defer file.Close()
		fmt.Printf("Todo list '%s' created successfully.\n", listName+".txt")
	} else if err != nil {
		return fmt.Errorf("error checking todo list file: %v", err)
	} else {
		fmt.Printf("Todo list '%s' already exists.\n", listName+".txt")
	}

	return nil
}
