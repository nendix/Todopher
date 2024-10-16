package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nendix/Todopher/internal/utils"
)

type AppState int

const (
	ViewTodos AppState = iota
	ViewFilteredTodos
	AddTodo
	EditTodo
	DeleteTodo
	MarkTodo
	UnmarkTodo
	SearchTodos
	SortTodos
)

type Model struct {
	todos      []utils.Todo
	filtered   []utils.Todo
	searchTerm string
	cursor     int
	state      AppState
	textInput  TextInput
	errMsg     string
	filePath   string // To store the todos file path
	sortType   string
}

// Init initializes the TUI model (required for tea.Model interface)
func (m Model) Init() tea.Cmd {
	todos, err := utils.ReadAllTodos(m.filePath)
	if err != nil {
		m.errMsg = "Error loading todos: " + err.Error()
	} else {
		m.todos = todos
	}
	return nil
}

// NewModel creates a new Model with default values
func NewModel(filePath string) Model {
	return Model{
		todos:     []utils.Todo{},
		cursor:    0,
		state:     ViewTodos,
		textInput: NewTextInput(),
		errMsg:    "",
		filePath:  filePath,
	}
}
