package tui

import (
	"fmt"
	"path/filepath"
	"strings"

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

		todoIDStr := todoIDStyle.Render(fmt.Sprintf("%03d", todo.ID))
		todoDesc := descriptionStyle.Render(todo.Description)
		todoDate := dateStyle.Render(todo.Date.String())

		s += fmt.Sprintf("%s %s %s %s - %s\n", cursor, todoIDStr, status, todoDesc, todoDate)
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
		todoDesc := todo.Description
		if m.searchTerm != "" {
			todoDesc = highlightSearchTerm(todoDesc, m.searchTerm)
		}

		todoIDStr := todoIDStyle.Render(fmt.Sprintf("%03d", todo.ID))
		todoDate := dateStyle.Render(todo.Date.String())

		s += fmt.Sprintf("%s %s %s %s - %s\n", cursor, todoIDStr, status, todoDesc, todoDate)
	}

	s += "\n(q)uit, (d)elete, (m)ark/unmark"
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
