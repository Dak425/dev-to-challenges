package pyramid

import "testing"

func TestPyramid(t *testing.T) {
	testCases := []struct {
		description string
		input       int
		expected    string
	}{
		{
			"small pyramid",
			3,
			"  *  \n *** \n*****\n",
		},
		{
			"no height",
			0,
			"",
		},
		{
			"negative height",
			-3,
			"",
		},
		{
			"big pyramid",
			7,
			"      *      \n     ***     \n    *****    \n   *******   \n  *********  \n *********** \n*************\n",
		},
	}

	for _, test := range testCases {
		if result := Pyramid(test.input); result != test.expected {
			t.Fatalf("FAIL: %s - Pyramid(%d): %s, expected: %s", test.description, test.input, result, test.expected)
		}
		t.Logf("PASS: %s", test.description)
	}
}
