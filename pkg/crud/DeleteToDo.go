package crud

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// DeleteToDo elimina un'attività dal file specificato
func DeleteToDo(filename string, id uint8) error {
	// Leggi il contenuto del file
	file, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Trova e rimuovi la riga corrispondente all'ID e aggiorna gli ID
	var updatedLines []string
	var newID uint8 = 1
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 4 {
			lineID, err := strconv.ParseUint(fields[0], 10, 8)
			if err == nil && uint8(lineID) == id {
				continue // Salta la riga da eliminare
			}
			fields[0] = fmt.Sprintf("%03d", newID) // Aggiorna l'ID
			newID++
		}
		updatedLines = append(updatedLines, strings.Join(fields, " "))
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
