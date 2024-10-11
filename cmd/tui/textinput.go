package tui

import tea "github.com/charmbracelet/bubbletea"

type TextInput struct {
	value  string
	cursor int
}

// NewTextInput initializes a new TextInput instance
func NewTextInput() TextInput {
	return TextInput{
		value:  "",
		cursor: 0,
	}
}

// NewTextInputWithValue initializes a new TextInput instance with a value
func NewTextInputWithValue(val string) TextInput {
	return TextInput{
		value:  val,
		cursor: len(val),
	}
}

// Update processes key events and updates the input state
func (ti *TextInput) Update(msg tea.KeyMsg) {
	switch msg.Type {
	case tea.KeyRunes:
		// Insert character at cursor position
		char := msg.String()
		ti.value = ti.value[:ti.cursor] + char + ti.value[ti.cursor:]
		ti.cursor += len(char)
	case tea.KeySpace:
		ti.value = ti.value[:ti.cursor] + " " + ti.value[ti.cursor:]
		ti.cursor++
	case tea.KeyBackspace, tea.KeyDelete:
		// Handle backspace/delete
		if ti.cursor > 0 {
			ti.value = ti.value[:ti.cursor-1] + ti.value[ti.cursor:]
			ti.cursor--
		}
	case tea.KeyLeft:
		if ti.cursor > 0 {
			ti.cursor--
		}
	case tea.KeyRight:
		if ti.cursor < len(ti.value) {
			ti.cursor++
		}
	}
}

// View renders the text input with the cursor
func (ti TextInput) View() string {
	cursor := "|"
	// Insert the cursor into the value at the correct position
	return ti.value[:ti.cursor] + cursor + ti.value[ti.cursor:]
}

// Value returns the current value of the input
func (ti TextInput) Value() string {
	return ti.value
}
