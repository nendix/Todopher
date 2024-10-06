package tgfuncs

import (
	"testing"

	"github.com/nendix/TaskGopher/internal/utils"
)

func TestParseToDo(t *testing.T) {
	tests := []struct {
		input       string
		expected    utils.ToDo
		expectError bool
	}{
		{
			input: "001 [ ] Buy groceries | 30/09/24",
			expected: utils.ToDo{
				ID:          1,
				Completed:   false,
				Description: "Buy groceries",
				Date:        utils.Date{Day: 30, Month: 9, Year: 24},
			},
			expectError: false,
		},
		{
			input: "002 [✓] Submit assignment | 15/10/24",
			expected: utils.ToDo{
				ID:          2,
				Completed:   true,
				Description: "Submit assignment",
				Date:        utils.Date{Day: 15, Month: 10, Year: 24},
			},
			expectError: false,
		},
		{
			input:       "003 [X] Invalid status | 20/10/24",
			expected:    utils.ToDo{},
			expectError: true,
		},
		{
			input:       "004 [ ] Missing due date",
			expected:    utils.ToDo{},
			expectError: true,
		},
	}

	for _, test := range tests {
		todo, err := utils.ParseToDo(test.input)
		if test.expectError {
			if err == nil {
				t.Errorf("Expected error for input '%s', but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for input '%s': %v", test.input, err)
			}
			if todo != test.expected {
				t.Errorf("For input '%s', expected %+v, but got %+v", test.input, test.expected, todo)
			}
		}
	}
}
