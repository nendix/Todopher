package crud

import (
	"bufio"
	"fmt"
	"os"
)

// AddToDo aggiunge un nuovo todo al file specificato
func AddToDo(filename, label, due string) error {
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
	_, err = fmt.Fprintf(file, "%03d [ ] %s per il %s\n", newID, label, due)
	if err != nil {
		return err
	}

	return nil
}

func readLastID(filename string) (uint8, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var lastID uint8
	isFirstLine := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if isFirstLine {
			// Salta la prima riga (intestazione)
			isFirstLine = false
			continue
		}
		var id uint8
		n, err := fmt.Sscanf(line, "%d ", &id)
		if err != nil || n != 1 {
			// Gestisci l'errore di parsing se necessario
			return 0, fmt.Errorf("error parsing ID from line: %s", line)
		}
		lastID = id
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lastID, nil
}
