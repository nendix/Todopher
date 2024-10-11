package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nendix/TaskGopher/internal/utils"
)

type AppState int

const (
	ViewTodos AppState = iota
	AddTodo
	EditTodo
	DeleteTodo
	MarkTodo
	UnmarkTodo
	SearchTodos
	SortTodos
)

type Model struct {
	todos     []utils.ToDo
	filtered  []utils.ToDo
	cursor    int
	state     AppState
	textInput TextInput
	errMsg    string
	filePath  string // To store the todos file path
	sortType  string
}

// Init initializes the TUI model (required for tea.Model interface)
func (m Model) Init() tea.Cmd {
	todos, err := utils.ReadAllToDos(m.filePath)
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
		todos:     []utils.ToDo{},
		cursor:    0,
		state:     ViewTodos,
		textInput: NewTextInput(),
		errMsg:    "",
		filePath:  filePath,
	}
}
