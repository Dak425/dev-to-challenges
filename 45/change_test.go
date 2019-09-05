package change

import "testing"

func equal(actual map[int]int, expected map[int]int) bool {
	if len(actual) != len(expected) {
		return false
	}

	for k, v := range actual {
		expectedVal, present := expected[k]

		if !present || v != expectedVal {
			return false
		}
	}

	return true
}

func TestChange(t *testing.T) {
	for _, test := range changeTests {
		if actual := Change(test.input); !equal(actual, test.expected) {
			t.Fatalf(
				"FAIL: %s - Change(%d): %v, expected: %v \n",
				test.description,
				test.input,
				actual,
				test.expected,
			)
		}
		t.Logf("PASS: %s \n", test.description)
	}
}
