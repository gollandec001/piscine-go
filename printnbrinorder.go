package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	digits := make([]int, 0)
	for n > 0 {
		digit := n % 10
		digits = append(digits, digit)
		n /= 10
	}
	sortSlice(digits)
	for _, digit := range digits {
		z01.PrintRune(rune(digit + '0'))
	}
}

func sortSlice(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}
