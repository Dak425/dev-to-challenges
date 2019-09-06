package likes

import (
	"testing"
)

var testCases = []struct {
	description string
	input       []string
	expected    string
}{
	{
		"no likes",
		[]string{},
		"no one likes this",
	},
	{
		"one like",
		[]string{"Mark"},
		"Mark likes this",
	},
	{
		"two likes",
		[]string{"Mark", "Jeff"},
		"Mark and Jeff like this",
	},
	{
		"three likes",
		[]string{"Mark", "Jeff", "Bob"},
		"Mark, Jeff, and Bob like this",
	},
	{
		"many likes",
		[]string{"Mark", "Jeff", "Bob", "Alice", "Susan"},
		"Mark, Jeff, and 3 others like this",
	},
}

func TestLikes(t *testing.T) {
	for _, test := range testCases {
		if result := Likes(test.input); result != test.expected {
			t.Fatalf("FAIL: %s - Likes(%v): %s, expected: %s \n", test.description, test.input, result, test.expected)
		}
		t.Logf("PASS: %s \n", test.description)
	}
}
