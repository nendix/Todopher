package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ShowMenu() {
	fmt.Println("\nTODO MENU: ")
	fmt.Println("1. Add todo")
	fmt.Println("2. Mark todo")
	fmt.Println("3. Unmark todo")
	fmt.Println("4. Delete todo")
	fmt.Println("5. Quit")
	fmt.Print("\nEnter your choice: ")
}

func ReadChoice() int {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	choice, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("Invalid input, please enter a number.")
		return 0
	}
	return choice
}

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

	// Verifica se il file è stato effettivamente creato
	created := false
	if info, err := file.Stat(); err == nil {
		if info.Size() == 0 {
			created = true
		}
	}

	return created, nil
}

func ReadString() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input[:len(input)-1] // Rimuovi il newline finale
}
