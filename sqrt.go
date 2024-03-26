package piscine

func Sqrt(nb int) int {
	gs := 1.0
	eps := 0.000001

	for i := 0; i < 1000; i++ {
		nextgs := 0.5 * (gs + float64(nb)/gs)
		if abs(nextgs-gs) < eps {
			break
		}
		gs = nextgs
	}
	if int(gs+0.5)*int(gs+0.5) == nb {
		return int(gs + 0.5)
	}
	return 0
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
