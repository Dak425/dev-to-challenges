package education

type Trip struct {
	Students []Student
}

func (t Trip) Richest() string {
	sCount := len(t.Students)
	if sCount == 0 {
		return "none"
	}
	if sCount == 1 {
		return t.Students[0].Name
	}
	poorest, richest := t.Students[0], t.Students[0]
	for i := 1; i < len(t.Students); i++ {
		if t.Students[i].Funds() < poorest.Funds() {
			poorest = t.Students[i]
		} else if t.Students[i].Funds() > richest.Funds() {
			richest = t.Students[i]
		}
	}
	if poorest.Funds() == richest.Funds() {
		return "all"
	}
	return richest.Name
}

type Student struct {
	Name     string
	Fives    int
	Tens     int
	Twenties int
}

func (s Student) Funds() int {
	return (s.Fives * 5) + (s.Tens * 10) + (s.Twenties * 20)
}
