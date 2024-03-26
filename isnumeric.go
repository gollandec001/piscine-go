package piscine

func IsNumeric(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if !isNumeric(c) {
			return false
		}
	}
	return true
}

func isNumeric(c rune) bool {
	return c >= '0' && c <= '9'
}
