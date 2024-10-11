package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.state {
		case ViewTodos:
			return m.updateViewTodos(msg)
		case AddTodo:
			return m.updateAddTodo(msg)
		case EditTodo:
			return m.updateEditTodo(msg)
		case SearchTodos:
			return m.updateSearchTodos(msg)
		case SortTodos:
			return m.updateSortTodos(msg)
		case DeleteTodo:
			return m.updateDeleteTodo(msg)
		}
	}
	return m, nil
}

func (m *Model) updateViewTodos(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "j":
		if m.cursor < len(m.todos)-1 {
			m.cursor++
		}
	case "a":
		m.state = AddTodo
		m.textInput = NewTextInput() // Reset the input for add mode
	case "e":
		m.state = EditTodo
		selectedTodo := m.todos[m.cursor]
		m.textInput = NewTextInputWithValue(selectedTodo.Description + " | " + selectedTodo.Date.String())
	case "d":
		m.state = DeleteTodo
	case "m":
		m.toggleMarkTodo()
		m.reloadTodos()
	case "s":
		m.state = SearchTodos
		m.textInput = NewTextInput() // Initialize the text input for searching
	case "o":
		m.state = SortTodos
	}
	return m, nil
}

func (m *Model) updateAddTodo(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyEnter:
		if m.textInput.Value() != "" {
			m.addTodo()
			m.state = ViewTodos
			m.textInput = NewTextInput() // Reset the input after adding
		}
	case tea.KeyEsc:
		m.state = ViewTodos
		m.textInput = NewTextInput() // Reset the input if cancelled
	default:
		m.textInput.Update(msg) // Update the text input based on key events
	}
	return m, nil
}

func (m *Model) updateEditTodo(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyEnter:
		if m.textInput.Value() != "" {
			m.editTodo()
			m.state = ViewTodos
		}
	case tea.KeyEsc:
		m.state = ViewTodos
	default:
		m.textInput.Update(msg) // Update the text input based on key events
	}
	return m, nil
}

func (m *Model) updateDeleteTodo(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyEnter:
		m.deleteTodo()
		m.state = ViewTodos
	case tea.KeyEsc:
		m.state = ViewTodos
	}
	return m, nil
}

func (m *Model) updateSearchTodos(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyEnter:
		searchTerm := m.textInput.Value()
		m.searchTodos(searchTerm)
		m.state = ViewTodos
	case tea.KeyEsc:
		m.resetSearch()
		m.state = ViewTodos
	default:
		m.textInput.Update(msg)
	}
	return m, nil
}

func (m *Model) updateSortTodos(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "d":
		m.sortTodos("date")
		m.state = ViewTodos
	case "s":
		m.sortTodos("status")
		m.state = ViewTodos
	case "esc":
		m.state = ViewTodos
	}
	return m, nil
}
