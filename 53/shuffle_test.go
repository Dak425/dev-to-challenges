package shuffle

import "testing"

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

func TestFaro(t *testing.T) {
	testCases := []struct {
		description string
		input       []string
		expected    []string
		expectErr   bool
	}{
		{
			"empty slice",
			[]string{},
			[]string{},
			false,
		},
		{
			"returns error on odd sized slice",
			[]string{"ace", "one", "two"},
			nil,
			true,
		},
		{
			"only two entries",
			[]string{"ace", "king"},
			[]string{"ace", "king"},
			false,
		},
		{
			"small deck",
			[]string{"ace", "two", "three", "four"},
			[]string{"ace", "three", "two", "four"},
			false,
		},
		{
			"post example",
			[]string{"ace", "two", "three", "four", "five", "six"},
			[]string{"ace", "four", "two", "five", "three", "six"},
			false,
		},
	}

	for _, test := range testCases {
		result, err := Faro(test.input)
		if err == nil && test.expectErr {
			t.Fatalf("FAIL: %s - Faro(%+v) - expected error to be thrown", test.description, test.input)
		}
		if !equal(result, test.expected) {
			t.Fatalf("FAIL: %s - Faro(%+v): %+v, %+v - expected '%+v'", test.description, test.input, result, err, test.expected)
		}
		t.Logf("PASS: %s", test.description)
	}
}
