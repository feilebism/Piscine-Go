package piscine

func IsPrintable(s string) bool {
	for _, printable := range s {
		if printable < 32 || printable > 126 {
			return false
		}
	}
	return true
}
