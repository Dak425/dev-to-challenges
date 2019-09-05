package ingots

const (
	lavaDuration     int = 800
	blazeRodDuration int = 120
	coalDuration     int = 80
	woodDuration     int = 15
)

// Requirements represents the differents types of fuel and in what amounts needed to complete the smelting job
type Requirements struct {
	Lava     int
	BlazeRod int
	Coal     int
	Wood     int
	Stick    int
}

// Fuel determines the fuel requirements to smelt the given number of ingots
func Fuel(ingots int) Requirements {
	if ingots <= 0 {
		return Requirements{0, 0, 0, 0, 0}
	}

	duration := ingots * 11

	return Requirements{
		Lava:     duration/lavaDuration + 1,
		BlazeRod: duration/blazeRodDuration + 1,
		Coal:     duration/coalDuration + 1,
		Wood:     duration/woodDuration + 1,
		Stick:    duration,
	}
}
