package tgfuncs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/nendix/TaskGopher/internal/utils"
)

// saveCurrentFile salva il nome della lista corrente nel file .env
func saveCurrentFile(listName string) error {
	// Ottieni il percorso della directory todo
	todoDir, err := utils.GetToDoDir()
	if err != nil {
		return fmt.Errorf("error getting todo directory: %v", err)
	}

	// Percorso del file .env
	envFilePath := filepath.Join(todoDir, ".env")

	// Carica le variabili d'ambiente dal file .env esistente
	envMap, err := godotenv.Read(envFilePath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to read .env file: %v", err)
	}

	// Aggiorna la variabile TODO_FILE con il nome della nuova lista
	envMap["TODO_FILE"] = listName

	// Scrivi nuovamente il file .env aggiornato
	err = godotenv.Write(envMap, envFilePath)
	if err != nil {
		return fmt.Errorf("failed to update .env file: %v", err)
	}

	return nil
}

// SetList verifica se il file della lista esiste in ~/todo/ o lo crea
func SetList(listName string) error {
	// Ottieni la directory ~/todo
	todoDir, err := utils.GetToDoDir()
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

	// Aggiorna il file .env con il nome del nuovo file della lista
	err = saveCurrentFile(listName + ".txt")
	if err != nil {
		return fmt.Errorf("failed to update .env file: %v", err)
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
		fmt.Printf("Todo list '%s' created successfully.\n", listName)
	} else if err != nil {
		return fmt.Errorf("error checking todo list file: %v", err)
	}
	return nil
}
