package population

// Threshold returns the number of years needed for the given population
// to reach a given population threshold
func Threshold(p0, aug, p int, percent float64) (year int) {
	for p0 < p {
		year++
		p0 += int(float64(p0)*(percent/100)) + aug
	}
	return
}
