package curly

// Validate will take a string of curly brackets, check it, and determine if it is valid (balanced and in proper order)
// "{}" = true
// "{{}" = false
// "{{}}" = true
func Validate(curlies string) bool {
	if curlies == "" {
		return true
	}

	balanced := 0

	for _, curly := range curlies {
		switch curly {
		case '{':
			balanced++
		case '}':
			balanced--
		}
		if balanced < 0 {
			return false
		}
	}

	return balanced == 0
}
