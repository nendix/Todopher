package tui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	cursorStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#fab387")) // Orange
	todoIDStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#89b4fa")) // Blue
	highlightStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#f38ba8")) // Red color
	descriptionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#f5e0dc")) // green
	dateStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#f5c2e7"))
)
