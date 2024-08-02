package tg

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// EditToDo modifica il testo e/o la data di scadenza di un'attività
func EditToDo(filename string, id uint8, newTodo string, newDate string) error {
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
		// Trova e aggiorna la riga corrispondente all'ID
		fields := strings.Fields(line)
		if len(fields) >= 4 {
			lineID, err := strconv.ParseUint(fields[0], 10, 8)
			if err == nil && uint8(lineID) == id {
				// Costruisci la nuova riga con il testo e/o la data di scadenza modificati
				fields[1] = "[ ]" // Rimuovi il marcatura esistente
				// La data di scadenza può essere nell'ultimo campo, ma assicurati di avere un formato corretto
				if len(fields) >= 4 {
					line = fmt.Sprintf("%03d %s %s for %s", id, fields[1], newTodo, newDate)
				}
			}
		}
		updatedLines = append(updatedLines, line)
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
