package tgfuncs

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// DeleteToDos elimina una lista di attività dal file specificato
func DeleteToDos(filePath string, ids []uint8) error {
	// Leggi il contenuto del file
	file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
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

	// Crea un set di ID da eliminare per una ricerca veloce
	idSet := make(map[uint8]struct{})
	for _, id := range ids {
		idSet[id] = struct{}{}
	}

	// Trova e rimuovi le righe corrispondenti agli ID e aggiorna gli ID
	var updatedLines []string
	var newID uint8 = 1
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 4 {
			lineID, err := strconv.ParseUint(fields[0], 10, 8)
			if err == nil {
				if _, exists := idSet[uint8(lineID)]; exists {
					continue // Salta le righe da eliminare
				}
				fields[0] = fmt.Sprintf("%03d", newID) // Aggiorna l'ID
				newID++
			}
		}
		updatedLines = append(updatedLines, strings.Join(fields, " "))
	}

	// Riscrivi il file con il contenuto aggiornato
	file, err = os.OpenFile(filePath, os.O_TRUNC|os.O_RDWR, 0666)
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
