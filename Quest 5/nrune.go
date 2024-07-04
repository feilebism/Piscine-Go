package piscine

func NRune(s string, n int) rune {
	result := []rune(s)
	length := 0
	for index := range result {
		length = index
	}

	if n-1 >= 0 && n-1 <= length {
		return result[n-1]
	}
	return 0
}
