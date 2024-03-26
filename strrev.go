package piscine

func StrRev(s string) string {
	runes := []rune(s)
	length := len(runes)
	left := 0
	right := length - 1
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
	return string(runes)
}
