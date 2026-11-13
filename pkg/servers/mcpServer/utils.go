package mcpServer

import (
	"fmt"
	"time"
)

func parseDOB(dobStr string) (time.Time, error) {
	const layout = "02/01/2006"

	return time.Parse(layout, dobStr)
}

func parseTOB(tobStr string) (time.Time, error) {
	const layout = "15:04"

	return time.Parse(layout, tobStr)
}

func CalculateAge(dobStr string) (int, error) {
	// Go uses the specific reference date Mon Jan 2 15:04:05 MST 2006 for layouts
	const layout = "02/01/2006"

	dob, err := parseDOB(dobStr)
	if err != nil {
		return 0, fmt.Errorf("invalid date format (expected dd/mm/yyyy): %w", err)
	}

	now := time.Now()
	if dob.After(now) {
		return 0, fmt.Errorf("date of birth cannot be in the future")
	}

	age := now.Year() - dob.Year()

	// Subtract 1 if the birthday hasn't occurred yet in the current year
	if now.Month() < dob.Month() || (now.Month() == dob.Month() && now.Day() < dob.Day()) {
		age--
	}

	return age, nil
}
