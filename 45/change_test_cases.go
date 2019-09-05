package change

var changeTests = []struct {
	description string
	input       int
	expected    map[int]int
}{
	{
		"no change",
		0,
		map[int]int{25: 0, 10: 0, 5: 0, 1: 0},
	},
	{
		"uses one coin",
		5,
		map[int]int{25: 0, 10: 0, 5: 1, 1: 0},
	},
	{
		"uses two coins",
		27,
		map[int]int{25: 1, 10: 0, 5: 0, 1: 2},
	},
	{
		"uses many coin",
		39,
		map[int]int{25: 1, 10: 1, 5: 0, 1: 4},
	},
}
