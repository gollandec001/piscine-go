package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	for frstdg := 0; frstdg <= 7; frstdg++ {
		for scnddg := frstdg + 1; scnddg <= 8; scnddg++ {
			for thrddg := scnddg + 1; thrddg <= 9; thrddg++ {
				if frstdg != scnddg && frstdg != thrddg && scnddg != thrddg {
					z01.PrintRune(rune('0' + frstdg))
					z01.PrintRune(rune('0' + scnddg))
					z01.PrintRune(rune('0' + thrddg))
					if frstdg != 7 || scnddg != 8 || thrddg != 9 {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
