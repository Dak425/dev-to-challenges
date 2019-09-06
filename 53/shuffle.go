package shuffle

import (
	"errors"
)

// Faro interweaves two halves of a slice, leaving the first and last elements unchanged
func Faro(deck []string) ([]string, error) {
	size := len(deck)

	if size%2 != 0 {
		return nil, errors.New("deck must be of an even size")
	}

	// if the size is empty or just two, simply return the deck
	if size <= 2 {
		return deck, nil
	}

	left := deck[:(size / 2)]
	right := deck[(size / 2):]

	shuffled := []string{}

	for i := range left {
		shuffled = append(shuffled, left[i], right[i])
	}

	return shuffled, nil
}
