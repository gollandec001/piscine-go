package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for za := 'z'; za >= 'a'; za-- {
		z01.PrintRune(za)
	}
	z01.PrintRune('\n')
}
