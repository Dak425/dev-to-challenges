package evolution

import "fmt"

const (
	negative string = "A negative evolution of %.f%%."
	positive string = "A positive evolution of %.f%%."
	neutral  string = "No evolution."
)

// Rate indicates the evolution percentage based on the before and after values
func Rate(before float64, after float64) string {
	if before == 0 {
		return fmt.Sprintf(positive, after*100)
	}

	if after == 0 {
		return fmt.Sprintf(negative, before*100)
	}

	rate := (after - before) / before * 100

	if rate > 0 {
		return fmt.Sprintf(positive, rate)
	}

	if rate < 0 {
		return fmt.Sprintf(negative, rate*-1)
	}

	return neutral
}
