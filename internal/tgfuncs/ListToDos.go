package tg

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nendix/TaskGopher/internal/utils"
)

// ListToDos stampa la lista di tutti i todo dal file specificato
func ListToDos(filePath string) error {
	listName, err := utils.ReadCurrentList()

	if listName != "" {
		fmt.Println("List:", listName)
	}
	file, err := os.Open(filePath)
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
