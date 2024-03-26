package piscine

func BasicAtoi2(s string) int {
	result := 0
	for _, char := range s {
		digit := int(char - '0')
		if digit < 0 || digit > 9 {
			return 0
		}
		result = result*10 + digit
	}
	return result
}
