package piscine

func TrimAtoi(s string) int {
	var result int
	sign := 1
	digitFound := false
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			digitFound = true
			result = result*10 + int(ch-'0')
		} else if ch == '-' && !digitFound {
			sign = -1
		}
	}
	return result * sign
}
