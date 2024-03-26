package piscine

func UltimateDivMod(a *int, b *int) {
	rtr := *a / *b
	gtr := *a % *b

	*a = rtr
	*b = gtr
}
