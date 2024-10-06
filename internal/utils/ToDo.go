package utils

import "fmt"

// ToDo represents a single todo item.
type ToDo struct {
	ID          uint8  // Unique identifier for the todo
	Completed   bool   // Completion status
	Description string // Description of the todo
	Date        Date   // Due date
}

// Date represents the date in DD/MM/YY format.
type Date struct {
	Day   int
	Month int
	Year  int
}

// String formats the Date into "DD/MM/YY".
func (d Date) String() string {
	return fmt.Sprintf("%02d/%02d/%04d", d.Day, d.Month, d.Year)
}

// String formats the ToDo into the file's string representation.
func (t ToDo) String() string {
	status := "[ ]"
	if t.Completed {
		status = "[✓]"
	}
	return fmt.Sprintf("%03d %s %s | %s", t.ID, status, t.Description, t.Date.String())
}
