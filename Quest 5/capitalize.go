package piscine

func Capitalize(s string) string {
	sr := []rune(s)
	prev := false
	for i := 0; i < len(sr); i++ {
		if !prev {
			if IsAlphaNumeric(sr[i]) {
				if sr[i] >= 97 && sr[i] <= 122 {
					sr[i] -= 32
				}
				prev = true
			}
		} else {
			if !IsAlphaNumeric(sr[i]) {
				prev = false
			} else {
				if sr[i] >= 65 && sr[i] <= 90 {
					sr[i] += 32
				}
			}
		}
	}
	return string(sr)
}

func IsAlphaNumeric(r rune) bool {
	if r >= 48 && r <= 57 {
		return true
	} else if r >= 65 && r <= 90 {
		return true
	} else if r >= 97 && r <= 122 {
		return true
	} else {
		return false
	}
}
