package body

import "testing"

type testCase struct {
	description string
	input       Attributes
	expected    string
}

func TestBMI(t *testing.T) {
	testCases := []testCase{
		{
			"underweight boi",
			Attributes{71, 116.0},
			underweight,
		},
		{
			"normal boi",
			Attributes{71, 147.0},
			normal,
		},
		{
			"overweight boi",
			Attributes{67, 190},
			overweight,
		},
		{
			"obese boi",
			Attributes{65, 210},
			obese,
		},
	}

	for _, test := range testCases {
		if result := BMI(test.input); result != test.expected {
			t.Fatalf("FAIL: %s - BMI(%+v): %s - expected %s", test.description, test.input, result, test.expected)
		}
		t.Logf("PASS: %s", test.description)
	}
}

func TestBMIMetric(t *testing.T) {
	testCases := []testCase{
		{
			"underweight boi",
			Attributes{1.8034, 52.61671},
			underweight,
		},
		{
			"normal boi",
			Attributes{1.8034, 66.67808},
			normal,
		},
		{
			"overweight boi",
			Attributes{1.7018, 86.18255},
			overweight,
		},
		{
			"obese boi",
			Attributes{1.651, 95.2544},
			obese,
		},
	}

	for _, test := range testCases {
		if result := BMIMetric(test.input); result != test.expected {
			t.Fatalf("FAIL: %s - BMIMetric(%+v): %s - expected %s", test.description, test.input, result, test.expected)
		}
		t.Logf("PASS: %s", test.description)
	}
}
