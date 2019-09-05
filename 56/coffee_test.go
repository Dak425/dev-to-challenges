package coffee

import "testing"

func TestPurchase(t *testing.T) {
	testCases := []struct {
		description string
		input       Order
		expected    string
	}{
		{
			"buy an Americano with exact change",
			Order{
				[]Product{
					Americano,
				},
				AmericanoPrice,
			},
			"Here is your Americano, have a nice day!",
		},
		{
			"buy an Americano without exact change",
			Order{
				[]Product{
					Americano,
				},
				2.30,
			},
			notExact,
		},
		{
			"buy a Latte with exact change",
			Order{
				[]Product{
					Latte,
				},
				LattePrice,
			},
			"Here is your Latte, have a nice day!",
		},
		{
			"buy a Latte without exact change",
			Order{
				[]Product{
					Latte,
				},
				5.00,
			},
			notExact,
		},
		{
			"buy a Flat White with exact change",
			Order{
				[]Product{
					FlatWhite,
				},
				FlatWhitePrice,
			},
			"Here is your Flat White, have a nice day!",
		},
		{
			"buy a Flat White without exact change",
			Order{
				[]Product{
					FlatWhite,
				},
				10.00,
			},
			notExact,
		},
		{
			"buy a Filter with exact change",
			Order{
				[]Product{
					Filter,
				},
				FilterPrice,
			},
			"Here is your Filter, have a nice day!",
		},
		{
			"buy a Filter without exact change",
			Order{
				[]Product{
					Filter,
				},
				4.00,
			},
			notExact,
		},
		{
			"buy two drinks with exact change",
			Order{
				[]Product{
					Americano,
					FlatWhite,
				},
				AmericanoPrice + FlatWhitePrice,
			},
			"Here is your Americano and Flat White, have a nice day!",
		},
		{
			"buy two drinks without exact change",
			Order{
				[]Product{
					Filter,
					Latte,
				},
				10.00,
			},
			notExact,
		},
		{
			"buy many drinks with exact change",
			Order{
				[]Product{
					Americano,
					Americano,
					FlatWhite,
					Latte,
					Filter,
				},
				(AmericanoPrice * 2) + FlatWhitePrice + LattePrice + FilterPrice,
			},
			"Here is your Americano, Americano, Flat White, Latte, and Filter, have a nice day!",
		},
		{
			"buy many drinks without exact change",
			Order{
				[]Product{
					Americano,
					Americano,
					FlatWhite,
					Latte,
					Filter,
				},
				30.00,
			},
			notExact,
		},
	}

	for _, test := range testCases {
		if result := Purchase(test.input); result != test.expected {
			t.Fatalf("FAIL: %s - Purchase(%+v): '%s' - expected '%s'", test.description, test.input, result, test.expected)
		}
		t.Logf("PASS: %s", test.description)
	}
}
