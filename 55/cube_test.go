package cube

import (
	"testing"
)

var testCases = []struct {
	description string
	input       int
	expected    int
}{
	{
		"challenge example 1",
		1071225,
		45,
	},
	{
		"challenge example 2",
		91716553919377,
		0,
	},
}

func TestCubes(t *testing.T) {
	for _, test := range testCases {
		if result := Cubes(test.input); result != test.expected {
			t.Fatalf("FAIL: %s - Cubes(%d): %d, expected %d", test.description, test.input, result, test.expected)
		}
		t.Logf("PASS: %s", test.description)
	}
}
