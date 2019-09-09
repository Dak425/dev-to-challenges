package josephus

import "math"

// Survivor indicates which man will be the survivor of a Josephus permutation
func Survivor(size uint, steps uint) uint {
	if size == 1 {
		return 0
	}

	if size < steps {
		return (Survivor(size-1, steps) + steps) % size
	}

	floor := size - uint(math.Floor(float64(size/steps)))

	return uint(math.Floor(float64((steps * ((Survivor(floor, steps) - size%steps) % floor)) / (steps - 1))))
}
