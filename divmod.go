package piscine

func DivMod(a int, b int, div *int, mod *int) {
	divide := a / b
	*div = divide
	*mod = a % b
}
