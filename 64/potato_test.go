package potato

import "testing"

type testCase struct {
	description string
	input       []uint
	expected    uint
	expectErr   bool
}

func TestDry(t *testing.T) {
	testCases := []testCase{
		{
			"example test",
			[]uint{99, 100, 98},
			50,
			false,
		},
		{
			"bigger change in water ratio",
			[]uint{99, 100, 95},
			20,
			false,
		},
		{
			"no change in water ratio",
			[]uint{87, 120, 87},
			120,
			false,
		},
		{
			"cannot use a zero mass of potatoes",
			[]uint{56, 0, 54},
			0,
			true,
		},
		{
			"cannot increase the water ratio",
			[]uint{42, 50, 50},
			0,
			true,
		},
	}
	for _, test := range testCases {
		result, err := Dry(test.input[0], test.input[1], test.input[2])
		if err == nil && test.expectErr {
			t.Fatalf("FAIL: %s - Dry(%d, %d, %d): expected error but nil was returned", test.description, test.input[0], test.input[1], test.input[2])
		}
		if result != test.expected {
			t.Fatalf("FAIL: %s - Dry(%d, %d, %d): %d - expected: %d", test.description, test.input[0], test.input[1], test.input[2], result, test.expected)
		}
		t.Logf("PASS: %s", test.description)
	}
}
