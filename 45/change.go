// Package change contains methods for interacting with US currency
package change

var coins = []int{25, 10, 5, 1}

// Change returns the fewest amount of each US coin to equal the amount of change indicated
func Change(amount int) map[int]int {
	usedCoins := map[int]int{25: 0, 10: 0, 5: 0, 1: 0}

	if amount == 0 {
		return usedCoins
	}

	remain := amount

	for _, coin := range coins {
		if remain == 0 {
			return usedCoins
		}

		mod := remain % coin
		if mod != remain {
			usedCoins[coin] = int(remain / coin)
			remain = mod
		}
	}

	return usedCoins
}
