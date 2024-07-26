package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFileIfNotExists(filename string) (bool, error) {
	// Ottieni la directory del file
	dir := filepath.Dir(filename)

	// Crea la directory se non esiste
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return false, fmt.Errorf("error creating directory: %w", err)
	}

	// Crea il file se non esiste
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return false, fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	// Controlla se il file è stato effettivamente creato
	fi, err := file.Stat()
	if err != nil {
		return false, err
	}

	// Controlla se il file era già presente
	return fi.Size() == 0, nil
}
