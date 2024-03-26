package piscine

func Atoi(s string) int {
	result := 0
	multiplier := 1
	for i, char := range s {
		if i == 0 && (char == '-' || char == '+') {
			if char == '-' {
				multiplier = -1
			}
			continue
		}
		if char < '0' || char > '9' {
			return 0
		}
		digit := int(char - '0')
		result = result*10 + digit
	}
	return result * multiplier
}
