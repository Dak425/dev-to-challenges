package potato

import "errors"

// Dry determines the weight of a mass of potatoes by reducing the current water ratio to the desired ratio
func Dry(p0, w0, p1 uint) (uint, error) {
	if p1 > p0 {
		return 0, errors.New("p1 cannot be greater than p0")
	}
	if w0 == 0 {
		return 0, errors.New("w0 cannot be zero")
	}
	if p0 == p1 {
		return w0, nil
	}
	return w0 * (100 - p0) / (100 - p1), nil
}
