package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	min := '0'
	max := '9'
	var st int = 0
	for i := min; i <= max; i++ {
		for j := min; j <= max; j++ {
			for k := i; k <= max; k++ {
				for l := min; l <= max; l++ {
					if i == k && (j > l || j == l) {
						continue
					}
					if st > 0 {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(k)
					z01.PrintRune(l)
					st = 1
				}
			}
		}
	}
	z01.PrintRune('\n')
}
