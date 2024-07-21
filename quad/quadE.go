package piscine

func QuadE(x, y int) {
	if x < 1 || y < 1 {
		return
	}
	topLine := CreateLine("A", "B", "C", x)
	midLine := CreateLine("B", " ", "B", x)
	botLine := CreateLine("C", "B", "A", x)
	FinalShape(topLine, midLine, botLine, y)
}
