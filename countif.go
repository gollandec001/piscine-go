package piscine

func CountIf(f func(string) bool, tab []string) int {
	cnt := 0
	for _, val := range tab {
		if f(val) {
			cnt++
		}
	}
	return cnt
}
