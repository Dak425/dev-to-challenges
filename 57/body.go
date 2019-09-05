package body

const (
	underweight string = "Underweight"
	normal      string = "Normal"
	overweight  string = "Overweight"
	obese       string = "Obese"
)

// Attributes represents some properties of a person's physique, such as weight and height
type Attributes struct {
	Height float64
	Weight float64
}

// BMI gives a description of your body health based on your attributes using body mass index
func BMI(attr Attributes) string {
	return status((attr.Weight / (attr.Height * attr.Height)) * 703)
}

// BMIMetric serves the same purpose as BMI but expects the attributes to be in the metric scale (kilograms, meters)
func BMIMetric(attr Attributes) string {
	return status(attr.Weight / (attr.Height * attr.Height))
}

func status(bmi float64) string {
	switch {
	case bmi <= 18.5:
		return underweight
	case bmi <= 25.0:
		return normal
	case bmi <= 30.0:
		return overweight
	default:
		return obese
	}
}
