package crud

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

// SortToDos ordina le attività nel file specificato per data o stato e riassegna gli ID
func SortToDos(filename, criteria string) error {
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

	// Funzione per ottenere la data di scadenza in formato time.Time
	parseDate := func(dateStr string) (time.Time, error) {
		const dateFormat = "02-01-06" // formato dd-mm-yy
		return time.Parse(dateFormat, dateStr)
	}

	// Funzione di confronto per l'ordinamento
	compare := func(i, j int) bool {
		// Estrai la data dalla fine della stringa
		lineI := lines[i]
		lineJ := lines[j]

		// Trova l'ultimo campo che dovrebbe essere la data
		fieldsI := strings.Fields(lineI)
		fieldsJ := strings.Fields(lineJ)

		// Assumiamo che la data sia sempre l'ultimo campo
		dateStrI := fieldsI[len(fieldsI)-1]
		dateStrJ := fieldsJ[len(fieldsJ)-1]

		// Parsing delle date
		dateI, errI := parseDate(dateStrI)
		dateJ, errJ := parseDate(dateStrJ)

		// Debug: Stampa gli errori di parsing se ci sono
		if errI != nil {
			fmt.Printf("Error parsing date for line %d: %v\n", i, errI)
		}
		if errJ != nil {
			fmt.Printf("Error parsing date for line %d: %v\n", j, errJ)
		}

		// Se ci sono errori nel parsing, non possiamo ordinare correttamente
		if errI != nil || errJ != nil {
			return false
		}

		// Ordinamento per data, mettendo la data più vicina prima
		return dateI.Before(dateJ)
	}

	// Ordinamento delle righe secondo il criterio
	sort.Slice(lines, compare)

	// Riassegna gli ID e prepara le righe aggiornate
	var updatedLines []string
	var newID uint8 = 1
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 1 {
			// Riassegna l'ID a partire da 001, 002, ...
			fields[0] = fmt.Sprintf("%03d", newID)
			newID++
			// Ricompone la riga con l'ID aggiornato
			updatedLine := strings.Join(fields, " ")
			updatedLines = append(updatedLines, updatedLine)
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
