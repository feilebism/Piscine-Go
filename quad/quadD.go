package piscine

func QuadD(x, y int) {
	if x < 1 || y < 1 {
		return
	}
	topBotLine := CreateLine("A", "B", "C", x)
	midLine := CreateLine("B", " ", "B", x)
	FinalShape(topBotLine, midLine, topBotLine, y)
}
