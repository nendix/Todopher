package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/nendix/TaskGopher/internal/utils"
)

func (m *Model) reloadTodos() {
	todos, err := utils.ReadAllToDos(m.filePath)
	if err != nil {
		m.errMsg = fmt.Sprintf("Error reloading todos: %v", err)
	} else {
		m.todos = todos
	}
}

var highlightStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#f38ba8")) // Red color
func highlightSearchTerm(description, searchTerm string) string {
	// Case-insensitive search
	lowerDesc := strings.ToLower(description)
	lowerTerm := strings.ToLower(searchTerm)

	// Find the index of the search term and replace all instances
	highlighted := description
	offset := 0
	for {
		idx := strings.Index(lowerDesc[offset:], lowerTerm)
		if idx == -1 {
			break
		}

		// Calculate the start and end index in the original description
		start := offset + idx
		end := start + len(searchTerm)

		// Apply the highlight style to the matched term
		highlighted = highlighted[:start] + highlightStyle.Render(description[start:end]) + highlighted[end:]

		// Move the offset forward
		offset = end
	}

	return highlighted
}
