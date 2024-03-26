package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	sortArgs(args)
	for _, arg := range args {
		printString(arg)
	}
}

func sortArgs(args []string) {
	for i := 0; i < len(args)-1; i++ {
		for j := i + 1; j < len(args); j++ {
			if compareArgs(args[i], args[j]) > 0 {
				args[i], args[j] = args[j], args[i]
			}
		}
	}
}

func compareArgs(arg1, arg2 string) int {
	for i := 0; i < len(arg1) && i < len(arg2); i++ {
		if arg1[i] != arg2[i] {
			return int(arg1[i]) - int(arg2[i])
		}
	}
	return len(arg1) - len(arg2)
}

func printString(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
