package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	switch m.state {
	case ViewTodos:
		return m.viewTodos()
	case AddTodo:
		return m.viewAddTodo()
	case EditTodo:
		return m.viewEditTodo()
	case DeleteTodo:
		return m.viewDeleteTodo()
	case SearchTodos:
		return m.viewSearchTodos()
	case SortTodos:
		return m.viewSortTodos()
	}
	return ""
}

var cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#89b4fa"))

func (m Model) viewTodos() string {
	if len(m.todos) == 0 {
		return "No todos found.\nPress 'a' to add a new todo."
	}

	s := "Todos:\n\n"
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

	s += "\n[q]uit, [a]dd, [e]dit, [d]elete, [m]ark/unmark, [s]earch, s[o]rt."
	return s
}

func (m Model) viewAddTodo() string {
	return fmt.Sprintf("Add a new todo:\n\n%s", m.textInput.View())
}

func (m Model) viewEditTodo() string {
	return fmt.Sprintf("Edit the selected todo:\n\n%s", m.textInput.View())
}

func (m Model) viewDeleteTodo() string {
	return fmt.Sprintf("Delete the selected todo? Press Enter to confirm, ESC to cancel.")
}

func (m Model) viewSearchTodos() string {
	return fmt.Sprintf("Filter: %s", m.textInput.View())
}

func (m Model) viewSortTodos() string {
	return "Criteria: [d]ate, [s]tatus"
}
