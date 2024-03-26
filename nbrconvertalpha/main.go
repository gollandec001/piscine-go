package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isWhitespace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r' || c == '\v' || c == '\f'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func atoi(s string) int {
	result := 0
	for _, ch := range []byte(s) {
		if !isDigit(ch) {
			return 0
		}
		result = result*10 + int(ch-'0')
	}
	return result
}

func printChar(n int, upper bool) {
	if n >= 1 && n <= 26 {
		letter := 'a' + rune(n-1)
		if upper {
			letter = 'A' + rune(n-1)
		}
		z01.PrintRune(letter)
	} else {
		z01.PrintRune(' ')
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}
	upper := false
	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}
	for _, arg := range args {
		num := atoi(arg)
		printChar(num, upper)
	}
	z01.PrintRune('\n')
}
