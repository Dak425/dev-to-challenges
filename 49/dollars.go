package dollars

import (
	"fmt"
)

// Dollars returns the string representation of a float as a dollar amount
func Dollars(amount float64) string {
	return fmt.Sprintf("$%.2f", amount)
}
