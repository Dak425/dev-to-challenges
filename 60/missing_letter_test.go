package find

import "testing"

type testCase struct {
	description string
	input       []rune
	expected    rune
}

func TestMissingLetter(t *testing.T) {
	testCases := []testCase{
		{
			"two characters",
			[]rune{'A', 'C'},
			'B',
		},
		{
			"dev-to example one",
			[]rune{'a', 'b', 'c', 'd', 'f'},
			'e',
		},
		{
			"dev-to example two",
			[]rune{'O', 'Q', 'R', 'S'},
			'P',
		},
	}

	for _, test := range testCases {
		if result := MissingLetter(test.input); result != test.expected {
			t.Fatalf("FAIL: %s - MissingLetter(%+v): %v - expected %v", test.description, test.input, result, test.expected)
		}
		t.Logf("PASS: %s", test.description)
	}
}
