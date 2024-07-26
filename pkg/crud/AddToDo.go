package crud

import (
	"bufio"
	"fmt"
	"os"
)

// AddToDo aggiunge un nuovo todo al file specificato
func AddToDo(filename, todo, date string) error {
	// Apri il file in modalità append e crea il file se non esiste
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	// Ottieni l'ID dell'ultimo todo
	lastID, err := readLastID(filename)
	if err != nil {
		return err
	}

	newID := lastID + 1

	// Scrivi il nuovo todo
	_, err = fmt.Fprintf(file, "%03d [ ] %s for %s\n", newID, todo, date)
	if err != nil {
		return err
	}

	return nil
}

// readLastID legge l'ultimo ID usato dal file
func readLastID(filename string) (uint8, error) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, nil // Se il file non esiste, ritorna 0
		}
		return 0, err
	}
	defer file.Close()

	var lastID uint8
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var id uint8
		_, err := fmt.Sscanf(line, "%d ", &id)
		if err == nil {
			lastID = id
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lastID, nil
}
