package century

import (
	"strconv"
)

// Century gives the text representation of what century the given date belongs to
func Century(date int) string {
	// if date is negative (BC), convert to positive
	if date < 0 {
		date *= -1
	}
	prefix := date/100 + 1

	var suffix string

	// if its one the weird teens centuries, suffix is "th"
	switch prefix % 100 {
	case 11, 12, 13:
		suffix = "th"
	default:
		switch prefix % 10 {
		case 1:
			suffix = "st"
		case 2:
			suffix = "nd"
		case 3:
			suffix = "rd"
		default:
			suffix = "th"
		}
	}

	return strconv.Itoa(prefix) + suffix
}
