package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	for position, firstrune := range s {
		if firstrune == FirstRune(toFind) {
			matches := 1
			for i := 2; i <= len(toFind); i++ {
				if NRune(toFind, i) == NRune(s, position+i) {
					matches++
				}
			}
			if matches == len(toFind) {
				return position
			}
		}
	}
	return -1
}
