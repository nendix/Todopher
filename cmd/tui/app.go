package tui

import (
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nendix/TaskGopher/internal/utils"
)

// StartTUI initializes and starts the TUI program
func StartTUI() error {
	// Initialize the model
	model, err := initializeModel()
	if err != nil {
		fmt.Println("Error initializing model:", err)
		return err
	}

	// Start the Bubble Tea program with the initial model
	p := tea.NewProgram(model)
	if err := p.Start(); err != nil {
		fmt.Println("Error running TUI program:", err)
		os.Exit(1)
	}

	return nil
}

// initializeModel sets up the initial state of the TUI model
func initializeModel() (Model, error) {
	todoDir, err := utils.GetToDoDir()
	if err != nil {
		fmt.Println("Error getting todo dir:", err)
		return Model{}, err
	}

	fileName, err := utils.ReadCurrentList()
	if err != nil {
		fmt.Println("Error reading current list:", err)
		return Model{}, err
	}
	filePath := filepath.Join(todoDir, fileName)

	// Load the todos from the file
	todos, err := utils.ReadAllToDos(filePath)
	if err != nil {
		return Model{
			filePath: filePath,
			errMsg:   fmt.Sprintf("Error loading todos: %v", err),
			state:    ViewTodos,
		}, nil // Return the model with an error message, but no actual error
	}

	// Return the initialized model
	return Model{
		todos:    todos,
		cursor:   0,
		state:    ViewTodos,
		filePath: filePath,
	}, nil
}
