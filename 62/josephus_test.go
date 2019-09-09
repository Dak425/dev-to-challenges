package josephus

import "testing"

type testCase struct {
	description string
	input       []uint
	expected    uint
}

func TestSurvivor(t *testing.T) {
	testCases := []testCase{
		{
			"1 man",
			[]uint{1, 5},
			0,
		},
		{
			"more men than steps",
			[]uint{7, 3},
			3,
		},
		{
			"less men than steps",
			[]uint{3, 5},
			0,
		},
		{
			"another with more men than steps",
			[]uint{15, 2},
			14,
		},
		{
			"another with less men than steps",
			[]uint{5, 17},
			3,
		},
	}

	for _, test := range testCases {
		if result := Survivor(test.input[0], test.input[1]); result != test.expected {
			t.Fatalf("FAIL: %s - Survivor(%+v): %d - expected: %d", test.description, test.input, result, test.expected)
		}
		t.Logf("PASS: %s", test.description)
	}
}
