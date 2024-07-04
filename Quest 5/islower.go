package piscine

func IsLower(s string) bool {
	letters := "abcdefghijklmnopqrstuvwxyz"
	counter := 0
	for _, a := range s {
		for _, b := range letters {
			if a == b {
				counter++
			}
		}
	}
	if counter == len(s) {
		return true
	} else {
		return false
	}
}
