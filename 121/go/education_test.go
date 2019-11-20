package education

import "testing"

type baseTestCase struct {
	description string
}

var bob = Student{"Bob", 4, 0, 0}
var mary = Student{"Mary", 0, 2, 0}
var john = Student{"John", 0, 0, 1}
var goat = Student{"Goat", 0, 0, 0}
var richie = Student{"Richie", 10, 20, 100}

func TestTrip_Richest(t *testing.T) {
	testCases := []struct {
		baseTestCase
		input    Trip
		expected string
	}{
		{
			baseTestCase{"no students"},
			Trip{},
			"all",
		},
		{
			baseTestCase{"one student"},
			Trip{[]Student{mary}},
			"all",
		},
		{
			baseTestCase{"many students with same amount of funds"},
			Trip{[]Student{mary, bob, john}},
			"all",
		},
		{
			baseTestCase{"many students with one having more"},
			Trip{[]Student{mary, bob, goat, richie}},
			richie.Name,
		},
	}
	for _, test := range testCases {
		if result := test.input.Richest(); result != test.expected {
			t.Fatalf("FAIL: %s - %+v.Funds(): '%s' - expected: '%s'", test.baseTestCase.description, test.input, result, test.expected)
		}
		t.Logf("PASS: %s", test.baseTestCase.description)
	}
}

func TestStudent_Funds(t *testing.T) {
	testCases := []struct {
		baseTestCase
		input    Student
		expected int
	}{
		{
			baseTestCase{"no moniez"},
			goat,
			0,
		},
		{
			baseTestCase{"only fives"},
			bob,
			20,
		},
		{
			baseTestCase{"only tens"},
			mary,
			20,
		},
		{
			baseTestCase{"only twenties"},
			john,
			20,
		},
		{
			baseTestCase{"all the moniez"},
			richie,
			2250,
		},
	}
	for _, test := range testCases {
		if result := test.input.Funds(); result != test.expected {
			t.Fatalf("FAIL: %s - %+v.Funds(): %d - expected: %d", test.baseTestCase.description, test.input, result, test.expected)
		}
		t.Logf("PASS: %s", test.baseTestCase.description)
	}
}
