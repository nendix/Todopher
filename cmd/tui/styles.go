package tui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	cursorStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#fab387")) // peach
	todoIDStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#89b4fa")) // blue
	highlightStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#f38ba8")) // red
	descriptionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#f5e0dc")) // rosewater
	dateStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#f5c2e7")) // pink
)
