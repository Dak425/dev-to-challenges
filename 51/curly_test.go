package curly

import "testing"

var testCases = []struct {
	description string
	input       string
	expected    bool
}{
	{
		"empty string",
		"",
		true,
	},
	{
		"simple string",
		"{}",
		true,
	},
	{
		"incorrect curlies",
		"{{}",
		false,
	},
	{
		"out of order curlies",
		"{}}{",
		false,
	},
	{
		"many curlies",
		"{}{{{}{}}}",
		true,
	},
}

func TestValidate(t *testing.T) {
	for _, test := range testCases {
		if result := Validate(test.input); result != test.expected {
			t.Fatalf("FAIL: %s, Validate(%s): %v - expected %v", test.description, test.input, result, test.expected)
		}
		t.Logf("PASS: %s", test.description)
	}
}
