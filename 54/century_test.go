package century

import "testing"

func TestCentury(t *testing.T) {
	testCases := []struct {
		description string
		input       int
		expected    string
	}{
		{
			"twenty third centry",
			2259,
			"23rd",
		},
		{
			"twelfth century",
			1124,
			"12th",
		},
		{
			"twenty first century",
			2000,
			"21st",
		},
		{
			"small centry",
			24,
			"1st",
		},
		{
			"odd centry name",
			1013,
			"11th",
		},
		{
			"large odd century name",
			11013,
			"111th",
		},
		{
			"negative century",
			-2000,
			"21st",
		},
	}

	for _, test := range testCases {
		if result := Century(test.input); result != test.expected {
			t.Fatalf("FAIL: %s - Centry(%d): %s - expected '%s'", test.description, test.input, result, test.expected)
		}
		t.Logf("PASS: %s", test.description)
	}
}
