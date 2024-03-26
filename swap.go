package piscine

func Swap(a *int, b *int) {
	po1 := *a
	po2 := *b
	*a = po2
	*b = po1
}
