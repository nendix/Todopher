package tui

import (
	"fmt"

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
