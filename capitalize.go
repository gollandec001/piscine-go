package piscine

func isAlphanumeric(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func toLower(c rune) rune {
	if c >= 'A' && c <= 'Z' {
		return c - 'A' + 'a'
	}
	return c
}

func toUpper(c rune) rune {
	if c >= 'a' && c <= 'z' {
		return c - 'a' + 'A'
	}
	return c
}

func Capitalize(s string) string {
	words := []rune(s)
	isFirst := true
	for i := 0; i < len(words); i++ {
		if !isAlphanumeric(words[i]) {
			isFirst = true
			continue
		}
		if isFirst {
			words[i] = toUpper(words[i])
			isFirst = false
		} else {
			words[i] = toLower(words[i])
		}
	}
	return string(words)
}
