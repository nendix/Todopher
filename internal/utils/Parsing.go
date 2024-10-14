package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ParseDate parses a date string in "DD/MM/YY" format into a Date struct.
func ParseDate(dateStr string) (Date, error) {
	var d Date
	_, err := fmt.Sscanf(dateStr, "%02d/%02d/%04d", &d.Day, &d.Month, &d.Year)
	if err != nil {
		return d, fmt.Errorf("invalid date format: %v", err)
	}
	return d, nil
}

// ParseTodo parses a todo line into a Todo struct using regex.
// Expected format: "ID [ ] Description - Due Date" or "ID [✓] Description | Due Date"
func ParseTodo(line string) (Todo, error) {
	var t Todo

	// Trim leading and trailing whitespace.
	line = strings.TrimSpace(line)

	// Define the regex pattern.
	pattern := `^(\d{3}) \[( |✓)\] (.+) - (\d{2}/\d{2}/\d{4})$`
	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(line)
	if len(matches) != 5 {
		return t, fmt.Errorf("invalid todo format")
	}

	// Extract matched groups.
	idStr := matches[1]
	status := matches[2]
	description := matches[3]
	dueDateStr := matches[4]

	// Parse ID.
	id, err := strconv.ParseUint(idStr, 10, 8)
	if err != nil {
		return t, fmt.Errorf("invalid ID '%s': %v", idStr, err)
	}
	t.ID = uint8(id)

	// Determine completion status.
	switch status {
	case " ":
		t.Completed = false
	case "✓":
		t.Completed = true
	default:
		return t, fmt.Errorf("unknown status: %s", status)
	}

	t.Description = description

	// Parse Due Date.
	t.Date, err = ParseDate(dueDateStr)
	if err != nil {
		return t, fmt.Errorf("invalid due date '%s': %v", dueDateStr, err)
	}

	return t, nil
}
