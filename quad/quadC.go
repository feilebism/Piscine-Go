package piscine

func QuadC(x, y int) {
	if x < 1 || y < 1 {
		return
	}
	topLine := CreateLine("A", "B", "A", x)
	midLine := CreateLine("B", " ", "B", x)
	botLine := CreateLine("C", "B", "C", x)
	FinalShape(topLine, midLine, botLine, y)
}
