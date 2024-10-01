package tgfuncs

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

// SortToDos sorts the todos in the specified file based on the given criteria ("date" or "status") and reassigns IDs.
func SortToDos(filePath, criteria string) error {
	// Step 1: Read all lines from the file.
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file for reading: %v", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		if len(line) < 7 {
			fmt.Printf("Skipping malformed line %d: %s\n", lineNumber, line)
			continue // Skip malformed lines.
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	// Step 2: Define date format consistent with input ("dd/mm/yy").
	const dateFormat = "02/01/06" // dd/mm/yy

	// Step 3: Sort based on criteria.
	switch strings.ToLower(criteria) {
	case "date":
		sort.Slice(lines, func(i, j int) bool {
			// Extract the due date from each line.
			dateStrI, errI := extractDueDate(lines[i])
			dateStrJ, errJ := extractDueDate(lines[j])

			if errI != nil {
				fmt.Printf("Error parsing date for line %d: %v\n", i+1, errI)
				return false // Keep original order if parsing fails.
			}
			if errJ != nil {
				fmt.Printf("Error parsing date for line %d: %v\n", j+1, errJ)
				return false
			}

			// Parse the dates.
			dateI, errI := time.Parse(dateFormat, dateStrI)
			dateJ, errJ := time.Parse(dateFormat, dateStrJ)

			if errI != nil || errJ != nil {
				// If either date fails to parse, keep original order.
				return false
			}

			// Sort by earlier dates first.
			return dateI.Before(dateJ)
		})
	case "status":
		sort.Slice(lines, func(i, j int) bool {
			// Extract the status from each line.
			statusI, errI := extractStatus(lines[i])
			statusJ, errJ := extractStatus(lines[j])

			if errI != nil || errJ != nil {
				// If either status fails to parse, keep original order.
				return false
			}

			// Define priority: incomplete ("[ ]") before complete ("[✓]")
			priority := map[string]int{
				"[ ]": 1,
				"[✓]": 2,
			}

			pi, okI := priority[statusI]
			pj, okJ := priority[statusJ]

			if !okI || !okJ {
				// If either status is unrecognized, keep original order.
				return false
			}

			// Sort by priority.
			return pi < pj
		})
	default:
		return fmt.Errorf("unknown sorting criteria: %s", criteria)
	}

	// Step 4: Reassign IDs and prepare updated lines.
	var updatedLines []string
	newID := 1
	for _, line := range lines {
		fields := strings.SplitN(line, " ", 3)
		if len(fields) < 3 {
			// Skip malformed lines.
			continue
		}

		// Reassign the ID with leading zeros (e.g., "001").
		newIDStr := fmt.Sprintf("%03d", newID)
		newLine := fmt.Sprintf("%s %s %s", newIDStr, fields[1], fields[2])
		updatedLines = append(updatedLines, newLine)
		newID++
	}

	// Step 5: Write updated lines back to the file.
	// Open the file for writing (truncate it).
	fileWrite, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return fmt.Errorf("failed to open file for writing: %v", err)
	}
	defer fileWrite.Close()

	writer := bufio.NewWriter(fileWrite)
	for _, line := range updatedLines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("failed to write to file: %v", err)
		}
	}

	// Flush the buffer to ensure all data is written.
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to flush data to file: %v", err)
	}

	return nil
}

// extractDueDate extracts the due date from a todo line.
// Expected format: "ID [ ] Task | Due Date"
func extractDueDate(line string) (string, error) {
	parts := strings.SplitN(line, " | ", 2)
	if len(parts) < 2 {
		return "", fmt.Errorf("missing due date")
	}
	return strings.TrimSpace(parts[1]), nil
}

// extractStatus extracts the status from a todo line.
// Expected format: "ID [ ] Task | Due Date"
func extractStatus(line string) (string, error) {
	fields := strings.Fields(line)
	if len(fields) < 2 {
		return "", fmt.Errorf("missing status")
	}
	status := strings.TrimSpace(fields[1])
	return status, nil
}
