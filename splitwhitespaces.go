package piscine

func SplitWhiteSpaces(s string) []string {
	words := []string{}
	currentWord := ""
	for _, ch := range s {
		if ch == ' ' || ch == '\t' || ch == '\n' {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(ch)
		}
	}
	if currentWord != "" {
		words = append(words, currentWord)
	}
	return words
}
