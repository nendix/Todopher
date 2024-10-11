package tui

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/nendix/TaskGopher/internal/utils"
)

func (m Model) View() string {
	switch m.state {
	case ViewTodos:
		return m.viewTodos()
	case ViewFilteredTodos:
		return m.viewFilteredTodos()
	case AddTodo:
		return m.viewAddTodo()
	case EditTodo:
		return m.viewEditTodo()
	case SearchTodos:
		return m.viewSearchTodos()
	case SortTodos:
		return m.viewSortTodos()
	}
	return ""
}

var cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#fab387"))

func (m Model) viewTodos() string {
	if len(m.todos) == 0 {
		return "No todos found.\nPress 'a' to add a new todo."
	}

	listName, err := utils.ReadCurrentList()
	if err != nil {
		return fmt.Sprintf("error reading current list: %v", err)
	}

	// Remove the .txt extension from the list name for display purposes.
	listName = strings.TrimSuffix(listName, filepath.Ext(listName))

	s := "List: " + listName + "\n\n"
	for i, todo := range m.todos {
		cursor := " "
		if m.cursor == i {
			cursor = cursorStyle.Render(">")
		}

		status := "[ ]"
		if todo.Completed {
			status = "[✓]"
		}

		s += fmt.Sprintf("%s %03d %s %s - %s\n", cursor, todo.ID, status, todo.Description, todo.Date.String())
	}

	s += "\n(q)uit, (a)dd, (e)dit, (d)elete, (m)ark/unmark, (s)earch, s(o)rt."
	return s
}

func (m Model) viewFilteredTodos() string {
	if len(m.filtered) == 0 {
		return "No todos match your search.\nPress 'esc' to return to the full list or 'q' to quit."
	}

	s := "Filtered Todos:\n\n"
	for i, todo := range m.filtered {
		cursor := " "
		if m.cursor == i {
			cursor = cursorStyle.Render(">")
		}

		status := "[ ]"
		if todo.Completed {
			status = "[✓]"
		}

		// Highlight the search term if it exists
		description := todo.Description
		if m.searchTerm != "" {
			description = highlightSearchTerm(description, m.searchTerm)
		}

		s += fmt.Sprintf("%s %03d %s %s - %s\n", cursor, todo.ID, status, description, todo.Date.String())
	}

	s += "\n(q)uit, (e)dit, (d)elete, (m)ark/unmark"
	return s
}

func (m Model) viewAddTodo() string {
	return fmt.Sprintf("Add a new todo:\n\n%s", m.textInput.View())
}

func (m Model) viewEditTodo() string {
	return fmt.Sprintf("Edit the selected todo:\n\n%s", m.textInput.View())
}

func (m Model) viewSearchTodos() string {
	return fmt.Sprintf("Filter: %s", m.textInput.View())
}

func (m Model) viewSortTodos() string {
	return "Criteria: (d)ate, (s)tatus"
}
