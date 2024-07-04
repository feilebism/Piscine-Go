package piscine

func IsAlpha(s string) bool {
	alphanumeric := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	counter := 0
	for _, a := range s {
		for _, b := range alphanumeric {
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
