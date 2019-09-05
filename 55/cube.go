package cube

// Cubes will determine the count of cubes necessary as n to construct a pile of volume m
// If no amount can equal m, 0 is returned
func Cubes(m int) int {
	var n, total int

	for total < m {
		n++
		total += (n * n * n)
	}

	if total == m {
		return n
	}

	return 0
}
