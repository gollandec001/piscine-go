package piscine

func LastRune(s string) rune {
	if len(s) > 0 {
		runes := []rune(s)
		return runes[len(runes)-1]
	}
	return 0
}
