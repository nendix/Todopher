package tg

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	todoDirName         = "todo"
	currentTodoFileName = "CURRENT_TODO_FILE.txt"
	initialFileName     = "todos.txt"
)

// Init crea la cartella ~/todo e il file ~/todo/CURRENT_TODO_FILE.txt con il contenuto todos.txt
func Init() error {
	// Ottieni la home directory dell'utente
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error getting home directory: %w", err)
	}

	// Costruisci il percorso della cartella todo e del file CURRENT_TODO_FILE.txt
	todoDir := filepath.Join(homeDir, todoDirName)
	currentTodoFilePath := filepath.Join(todoDir, currentTodoFileName)

	// Debug: Stampa il percorso della directory e verifica se esiste
	fmt.Printf("Attempting to create directory at: %s\n", todoDir)

	// Controlla se esiste già una directory o un file con il nome todoDir
	info, err := os.Stat(todoDir)
	if err == nil {
		if !info.IsDir() {
			return fmt.Errorf("file exists with the name '%s', but it's not a directory", todoDir)
		}
		fmt.Println("Directory already exists.")
	} else if os.IsNotExist(err) {
		// Se non esiste, crea la cartella todo
		err = os.MkdirAll(todoDir, 0755)
		if err != nil {
			return fmt.Errorf("error creating todo directory: %w", err)
		}
		fmt.Println("Directory created successfully.")
	} else {
		return fmt.Errorf("error checking todo directory: %w", err)
	}

	// Crea il file CURRENT_TODO_FILE.txt se non esiste
	_, err = os.Stat(currentTodoFilePath)
	if os.IsNotExist(err) {
		file, err := os.Create(currentTodoFilePath)
		if err != nil {
			return fmt.Errorf("error creating current todo file: %w", err)
		}
		defer file.Close()

		// Scrivi il nome del file iniziale todos.txt nel file CURRENT_TODO_FILE.txt
		_, err = file.WriteString(initialFileName)
		if err != nil {
			return fmt.Errorf("error writing to current todo file: %w", err)
		}
		fmt.Println("File created and initialized successfully.")
	} else if err != nil {
		return fmt.Errorf("error checking current todo file: %w", err)
	}

	return nil // Se tutto è andato bene, non restituisce errore
}
