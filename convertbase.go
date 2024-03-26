package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	decimalValue := ConvertToDecimal(nbr, baseFrom)
	return ConvertFromDecimal(decimalValue, baseTo)
}

func ConvertToDecimal(nbr, baseFrom string) int {
	base := len(baseFrom)
	decimalValue := 0
	power := 1
	for i := len(nbr) - 1; i >= 0; i-- {
		digit := getIndex(baseFrom, nbr[i])
		decimalValue += digit * power
		power *= base
	}
	return decimalValue
}

func ConvertFromDecimal(decimalValue int, baseTo string) string {
	base := len(baseTo)
	result := ""
	for decimalValue > 0 {
		remainder := decimalValue % base
		result = string(baseTo[remainder]) + result
		decimalValue /= base
	}
	return result
}

func getIndex(base string, ch byte) int {
	for i := 0; i < len(base); i++ {
		if base[i] == ch {
			return i
		}
	}
	return -1
}
