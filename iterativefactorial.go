package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	des := 1

	for i := 1; i <= nb; i++ {
		prevdes := des

		des *= i

		if des < prevdes {
			return 0
		}
	}
	return des
}
