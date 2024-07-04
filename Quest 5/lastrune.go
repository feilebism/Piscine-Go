package piscine

func LastRune(s string) rune {
	letter := []rune(s)
	return letter[len(letter)-1]
}
