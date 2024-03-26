package piscine

func IsPrintable(s string) bool {
	for _, c := range s {
		if !isPrintable(c) {
			return false
		}
	}
	return true
}

func isPrintable(c rune) bool {
	return c >= 32 && c <= 126
}
