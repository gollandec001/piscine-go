package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for dg := '0'; dg <= '9'; dg++ {
		z01.PrintRune(dg)
	}
	z01.PrintRune('\n')
}
