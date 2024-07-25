package crud

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// MarkToDo segna un'attività come completata nel file specificato
func MarkToDo(filename string, id uint8) error {
	// Leggi il contenuto del file
	file, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	// Leggi tutte le righe del file
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	// Trova e aggiorna la riga corrispondente all'ID
	var updatedLines []string
	for i, line := range lines {
		if i == 0 {
			// Mantieni l'intestazione
			updatedLines = append(updatedLines, line)
			continue
		}

		// Parsing della riga
		parts := strings.Fields(line)
		if len(parts) >= 4 {
			lineID, err := strconv.ParseUint(parts[0], 10, 8)
			if err == nil && uint8(lineID) == id {
				parts[1] = "[X]" // Segna come completato
			}
			updatedLines = append(updatedLines, strings.Join(parts, " "))
		} else {
			updatedLines = append(updatedLines, line)
		}
	}

	// Riscrivi il file con il contenuto aggiornato
	file, err = os.OpenFile(filename, os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range updatedLines {
		if _, err := fmt.Fprintln(file, line); err != nil {
			return err
		}
	}

	return nil
}
