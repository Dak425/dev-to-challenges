package population

import "testing"

type testCase struct {
	description string
	p0          int
	aug         int
	p           int
	percent     float64
	expected    int
}

func TestThreshold(t *testing.T) {
	testCases := []testCase{
		testCase{"first example", 1500, 100, 5000, 5, 15},
		testCase{"second example", 1500000, 10000, 2000000, 2.5, 10},
	}

	for _, test := range testCases {
		if result := Threshold(test.p0, test.aug, test.p, test.percent); result != test.expected {
			t.Fatalf("FAIL: %s - Threshold(%d, %d, %d, %f): %d - expected: %d", test.description, test.p0, test.aug, test.p, test.percent, result, test.expected)
		}
		t.Logf("PASS: %s", test.description)
	}
}
