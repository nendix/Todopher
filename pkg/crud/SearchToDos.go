package crud

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// SearchToDos cerca nel file specificato le righe che contengono la parola chiave
func SearchToDos(filename, keyword string) error {
	// Leggi il contenuto del file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Cerca la parola chiave in ogni riga del file
	scanner := bufio.NewScanner(file)
	found := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			fmt.Println(line)
			found = true
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if !found {
		return errors.New("No todos found with the given keyword.")
	}

	return nil
}
