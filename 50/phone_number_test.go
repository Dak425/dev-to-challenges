package phonenumbers

import "testing"

var testCases = []struct {
	description string
	input       string
	expected    []string
}{
	{
		"empty string",
		"",
		[]string{},
	},
	{
		"single digit",
		"5",
		[]string{"2", "4", "6", "8"},
	},
	{
		"many digits",
		"4531",
		[]string{
			"1531",
			"5531",
			"7531",
			"4231",
			"4431",
			"4631",
			"4831",
			"4521",
			"4561",
			"4532",
			"4534",
		},
	},
}

func equal(result []string, expected []string) bool {
	if len(result) != len(expected) {
		return false
	}

	for i := range result {
		if result[i] != expected[i] {
			return false
		}
	}

	return true
}

func TestPhoneNeighbors(t *testing.T) {
	for _, test := range testCases {
		if result := PhoneNeighbors(test.input); !equal(result, test.expected) {
			t.Fatalf("FAIL: %s - PhoneNeighbors(%s): %v - expected: %v", test.description, test.input, result, test.expected)
		}
		t.Logf("PASS: %s", test.description)
	}
}
