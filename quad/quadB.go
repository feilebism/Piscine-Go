package piscine

func QuadB(x, y int) {
	if x < 1 || y < 1 {
		return
	}
	topLine := CreateLine("/", "*", "\\", x)
	midLine := CreateLine("*", " ", "*", x)
	botLine := CreateLine("\\", "*", "/", x)
	FinalShape(topLine, midLine, botLine, y)
}
