package piscine

func Abort(a, b, c, d, e int) int {
	table := []int{a, b, c, d, e}
	n := len(table)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}

	return table[2]
}
