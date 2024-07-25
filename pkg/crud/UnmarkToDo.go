package crud

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// MarkToDo segna un'attività come completata nel file specificato
func UnmarkToDo(filename string, id uint8) error {
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

	// Trova e aggiorna la riga corrispondente all'ID
	var updatedLines []string
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 4 {
			lineID, err := strconv.ParseUint(fields[0], 10, 8)
			if err == nil && uint8(lineID) == id {
				fields[1] = "[ ]" // Segna come completato
			}
			updatedLines = append(updatedLines, strings.Join(fields, " "))
		}
	}

	// Scrivi il contenuto aggiornato nel file
	file, err = os.OpenFile(filename, os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range updatedLines {
		_, err := fmt.Fprintln(file, line)
		if err != nil {
			return err
		}
	}

	return nil
}
