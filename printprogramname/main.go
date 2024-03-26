package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) > 0 {
		name := args[0][2:]
		for _, ch := range name {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
