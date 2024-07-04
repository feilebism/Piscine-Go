package piscine

func IsUpper(s string) bool {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	counter := 0
	for _, str := range s {
		for _, upperletter := range letters {
			if str == upperletter {
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
