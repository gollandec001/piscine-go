package piscine

func BasicAtoi(s string) int {
	r := []rune(s)
	result := 0
	for _, char := range r {
		digit := int(char - '0')
		result = result*10 + digit
	}
	return result
}
