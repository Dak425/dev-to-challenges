package evolution

import "testing"

type testCase struct {
	description string
	input       []float64
	expected    string
}

func TestRate(t *testing.T) {
	testCases := []testCase{
		{
			"a positive evolution rate",
			[]float64{11.29, 45.79},
			"A positive evolution of 306%.",
		},
		{
			"a negative evolution rate",
			[]float64{95.12, 66.84},
			"A negative evolution of 30%.",
		},
		{
			"zero before value",
			[]float64{0, 27.35},
			"A positive evolution of 2735%.",
		},
		{
			"zero after value",
			[]float64{41.26, 0},
			"A negative evolution of 4126%.",
		},
		{
			"same before and after values",
			[]float64{1.26, 1.26},
			"No evolution.",
		},
	}

	for _, test := range testCases {
		if result := Rate(test.input[0], test.input[1]); result != test.expected {
			t.Fatalf("FAIL: %s - Rate(%+v): %s - expected: %s", test.description, test.input, result, test.expected)
		}
		t.Logf("PASS: %s", test.description)
	}
}
