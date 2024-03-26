package piscine

func AlphaCount(s string) int {
	cnt := 0
	for _, chr := range s {
		if (chr >= 'a' && chr <= 'z') || (chr >= 'A' && chr <= 'Z') {
			cnt++
		}
	}
	return cnt
}
