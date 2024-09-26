package tg

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// UnmarkToDos deseleziona più attività nel file specificato
func MarkToDos(filePath string, ids []uint8) error {
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

	// Crea un set di ID per una verifica rapida
	idSet := make(map[uint8]struct{})
	for _, id := range ids {
		idSet[id] = struct{}{}
	}

	// Trova e aggiorna le righe corrispondenti agli ID
	var updatedLines []string
	for _, line := range lines {
		lineIDStr := strings.Fields(line)[0]
		lineID, err := strconv.ParseUint(lineIDStr, 10, 8)
		if err != nil {
			updatedLines = append(updatedLines, line)
			continue
		}

		// Controlla se l'ID è nella lista degli ID da deselezionare
		if _, found := idSet[uint8(lineID)]; found {
			// Sostituisci [X] con [ ]
			line = strings.Replace(line, "[ ]", "[✓]", 1)
		}

		updatedLines = append(updatedLines, line)
	}

	// Scrivi il contenuto aggiornato nel file
	file, err = os.OpenFile(filePath, os.O_TRUNC|os.O_RDWR, 0666)
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
