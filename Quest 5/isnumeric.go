package piscine

func IsNumeric(s string) bool {
	numbers := "0123456789"
	counter := 0
	for _, a := range s {
		for _, b := range numbers {
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
