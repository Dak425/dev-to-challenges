package ingots

import "testing"

func TestFuel(t *testing.T) {
	testCases := []struct {
		description string
		input       int
		expected    Requirements
	}{
		{
			"a few ingots",
			2,
			Requirements{1, 1, 1, 2, 22},
		},
		{
			"wow that is alot of ingots",
			500,
			Requirements{7, 46, 69, 367, 5500},
		},
		{
			"negative amount of ingots",
			-5,
			Requirements{0, 0, 0, 0, 0},
		},
		{
			"no ingots",
			0,
			Requirements{0, 0, 0, 0, 0},
		},
	}

	for _, test := range testCases {
		if result := Fuel(test.input); result != test.expected {
			t.Fatalf("FAIL: %s - Fuel(%d): %+v - expected: %+v", test.description, test.input, result, test.expected)
		}
		t.Logf("PASS: %s", test.description)
	}
}
