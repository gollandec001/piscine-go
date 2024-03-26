package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}
	if nbr < 0 {
		z01.PrintRune('-')
		if nbr == -9223372036854775808 {
			printMinInt(base)
			return
		}
		nbr = -nbr
	}
	var result string
	length := len(base)
	for nbr > 0 {
		result = string(base[nbr%length]) + result
		nbr /= length
	}
	printStr(result)
}

func isValidBase(base string) bool {
	length := len(base)
	if length < 2 {
		return false
	}
	if containsChar(base, '+') || containsChar(base, '-') {
		return false
	}
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}
	return true
}

func containsChar(s string, c byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return true
		}
	}
	return false
}

func printStr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
}

func printMinInt(base string) {
	length := len(base)
	dividend := -9223372036854775808
	var result string
	for dividend <= -length {
		remainder := dividend % length
		result = string(base[-remainder]) + result
		dividend /= length
	}
	result = string(base[-dividend]) + result
	printStr(result)
}
