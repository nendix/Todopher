package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFileIfNotExists(filename string) (bool, error) {
	// Ottieni la directory del file
	dir := filepath.Dir(filename)

	// Crea la directory se non esiste
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return false, fmt.Errorf("error creating directory: %w", err)
	}

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

func PrintHelp() {
	fmt.Println("Usage: todo [command] [options]")
	fmt.Println("Commands:")
	fmt.Println("  list                                     - List all todos")
	fmt.Println("  add [todo] [dd-mm-yy]                    - Add a new todo")
	fmt.Println("  edit [id] [new todo] [new dd-mm-yy]      - Edit a todo")
	fmt.Println("  mark [id1 id2 ...]                       - Mark todos as completed")
	fmt.Println("  unmark [id1 id2 ...]                     - Unmark todos as not completed")
	fmt.Println("  search [key_word]                        - List all todos that contain the keyword")
	fmt.Println("  sort [id] [by_status|by_date]            - Sort todos by status or by date")
	fmt.Println("  delete [id1 id2 ...]                     - Delete todos")
}
