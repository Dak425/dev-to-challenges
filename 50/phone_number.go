package phonenumbers

// PhoneNeighbors returns all the phone numbers that are adjacent to the given phone number
func PhoneNeighbors(phone string) []string {
	numbers := []string{}

	if phone == "" {
		return numbers
	}

	runes := []rune(phone)

	for i, r := range runes {
		// Copy the phone number so we can replace the current digit with its neighbors during the loop
		instance := append(runes[:0:0], runes...)

		// Get the neighbors of this digit
		n := neighbors(r)

		// Loop over the neighbors
		for _, num := range n {
			// Replace the digit at this index with the neighbor and append the result to the slice
			instance[i] = num
			numbers = append(numbers, string(instance))
		}
	}

	return numbers
}

func neighbors(number rune) []rune {
	switch number {
	case '0':
		return []rune{'8'}
	case '1':
		return []rune{'2', '4'}
	case '2':
		return []rune{'1', '3', '5'}
	case '3':
		return []rune{'2', '6'}
	case '4':
		return []rune{'1', '5', '7'}
	case '5':
		return []rune{'2', '4', '6', '8'}
	case '6':
		return []rune{'3', '5', '9'}
	case '7':
		return []rune{'4', '8'}
	case '8':
		return []rune{'0', '5', '7', '9'}
	case '9':
		return []rune{'6', '8'}
	default:
		return []rune{}
	}
}
