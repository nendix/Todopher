package tui

import (
	"fmt"
	"strings"

	"github.com/nendix/TaskGopher/internal/tgfuncs"
	"github.com/nendix/TaskGopher/internal/utils"
)

func (m *Model) addTodo() {
	// Get the value from the text input
	inputValue := m.textInput.Value()

	// Ensure input is in the format: "Description | DD/MM/YYYY"
	parts := strings.SplitN(inputValue, "|", 2)
	if len(parts) != 2 {
		m.errMsg = "Invalid format. Use: Description | DD/MM/YYYY"
		return
	}

	// Trim and parse the input parts
	description := strings.TrimSpace(parts[0])
	dueDateStr := strings.TrimSpace(parts[1])

	// Validate the due date format
	_, err := utils.ParseDate(dueDateStr)
	if err != nil {
		m.errMsg = fmt.Sprintf("Invalid date format: %v. Use DD/MM/YYYY.", err)
		fmt.Println(m.errMsg) // Debug log
		return
	}

	// Call the AddToDo function to add the todo
	err = tgfuncs.AddToDo(m.filePath, description, dueDateStr)
	if err != nil {
		m.errMsg = fmt.Sprintf("Error adding todo: %v", err)
	} else {
		fmt.Println("Todo added successfully!") // Debug message
		m.reloadTodos()                         // Reload the todos to reflect the new addition
	}
}

func (m *Model) editTodo() {
	// Get the value from the text input
	inputValue := m.textInput.Value()

	// Ensure input is in the format: "Description | DD/MM/YYYY"
	parts := strings.SplitN(inputValue, "|", 2)
	if len(parts) != 2 {
		m.errMsg = "Invalid format. Use: Description | DD/MM/YYYY"
		return
	}

	// Trim and parse the input parts
	description := strings.TrimSpace(parts[0])
	dueDateStr := strings.TrimSpace(parts[1])

	// Validate the due date format
	_, err := utils.ParseDate(dueDateStr)
	if err != nil {
		m.errMsg = fmt.Sprintf("Invalid date format: %v. Use DD/MM/YYYY.", err)
		return
	}

	// Call the EditToDo function to edit the todo
	err = tgfuncs.EditToDo(m.filePath, m.todos[m.cursor].ID, description, dueDateStr)
	if err != nil {
		m.errMsg = fmt.Sprintf("Error editing todo: %v", err)
	} else {
		m.reloadTodos()
	}
}

func (m *Model) deleteTodo() {
	err := tgfuncs.DeleteToDos(m.filePath, []uint8{m.todos[m.cursor].ID})
	if err != nil {
		m.errMsg = fmt.Sprintf("Error deleting todo: %v", err)
	} else {
		m.reloadTodos()
	}
}

func (m *Model) toggleMarkTodo() {
	todo := m.todos[m.cursor]

	// Check if the todo is already marked as completed
	var err error
	if todo.Completed {
		// If completed, call UnmarkToDo to unmark it
		err = tgfuncs.UnmarkToDos(m.filePath, []uint8{todo.ID})
	} else {
		// If not completed, call MarkToDos to mark it as completed
		err = tgfuncs.MarkToDos(m.filePath, []uint8{todo.ID})
	}

	// Handle errors if any
	if err != nil {
		m.errMsg = fmt.Sprintf("Error marking/unmarking todo: %v", err)
	} else {
		m.reloadTodos() // Reload the todos to update the view
	}
}

func (m *Model) searchTodos(searchTerm string) {
	var filtered []utils.ToDo
	for _, todo := range m.todos {
		if strings.Contains(strings.ToLower(todo.Description), strings.ToLower(searchTerm)) {
			filtered = append(filtered, todo)
		}
	}
	m.filtered = filtered
}

// resetSearch clears the search filter
func (m *Model) resetSearch() {
	m.filtered = m.todos
}

func (m *Model) sortTodos(criteria string) {
	err := tgfuncs.SortToDos(m.filePath, criteria)
	if err != nil {
		m.errMsg = fmt.Sprintf("Error sorting todos: %v", err)
		return
	}

	// Reload the todos to reflect the sorted order
	todos, err := utils.ReadAllToDos(m.filePath)
	if err != nil {
		m.errMsg = fmt.Sprintf("Error reloading todos: %v", err)
		return
	}

	// Update the model's todos and filtered list
	m.todos = todos
	m.filtered = todos
}
