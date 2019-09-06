package pyramid

// Pyramid creates a text pyramid of the given height
func Pyramid(height int) string {
	if height <= 0 {
		return ""
	}

	// Determine width of pyramid base
	width := (height * 2) - 1

	// Get middle of pyramid
	middle := int(width / 2)
	row := 0

	// Initialize a base row with only spaces
	base := make([]rune, width, width)
	for i := range base {
		base[i] = ' '
	}

	// Set the top row 'block'
	base[middle] = '*'

	var pyramid string

	// Build the pyramid starting from the top
	for row < height {
		base[middle+row] = '*'
		base[middle-row] = '*'
		pyramid += string(base) + "\n"

		row++
	}

	return pyramid
}
