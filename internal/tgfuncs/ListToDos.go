package tg

import (
	"bufio"
	"fmt"
	"os"
)

// ListToDos stampa la lista di tutti i todo dal file specificato
func ListToDos(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
