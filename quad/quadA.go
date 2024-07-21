package piscine

import "fmt"

func QuadA(x, y int) {
	if x < 1 || y < 1 {
		return
	}
	topBotLine := CreateLine("o", "-", "o", x)
	midLine := CreateLine("|", " ", "|", x)
	FinalShape(topBotLine, midLine, topBotLine, y)
}

func CreateLine(a, b, c string, width int) string {
	NewStr := ""
	if width == 1 {
		NewStr = a
	} else {
		NewStr += a
		for i := 0; i < width-2; i++ {
			NewStr += b
		}
		NewStr += c
	}
	return NewStr
}

func FinalShape(top, mid, bot string, height int) {
	if height == 1 {
		fmt.Println(top)
	} else {
		fmt.Println(top)
		for i := 0; i < height-2; i++ {
			fmt.Println(mid)
		}
		fmt.Println(bot)
	}
}
