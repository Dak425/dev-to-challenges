package dollars

import (
	"testing"
)

var testCases = []struct {
	description string
	input       float64
	expected    string
}{
	{
		"only cents",
		.04,
		"$0.04",
	},
	{
		"only dollars",
		3,
		"$3.00",
	},
	{
		"cents & dollars",
		3.14,
		"$3.14",
	},
	{
		"amounts with more than two decimal places round to the nearest cent",
		4.478,
		"$4.48",
	},
}

func TestDollars(t *testing.T) {
	for _, test := range testCases {
		if result := Dollars(test.input); result != test.expected {
			t.Fatalf("FAILED: %s - Dollars(%f): %s, expected %s \n", test.description, test.input, result, test.expected)
		}
		t.Logf("PASS: %s \n", test.description)
	}
}
