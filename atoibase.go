package piscine

func AtoiBase(s string, base string) int {
	if len(base) < 2 || containsSign(base) || !isUnique(base) || !isValidString(s, base) {
		return 0
	}
	baseMap := make(map[byte]int)
	for i := 0; i < len(base); i++ {
		baseMap[base[i]] = i
	}
	result := 0
	for i := 0; i < len(s); i++ {
		result = result*len(base) + baseMap[s[i]]
	}
	return result
}

func containsSign(base string) bool {
	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			return true
		}
	}
	return false
}

func isUnique(base string) bool {
	charSet := make(map[byte]bool)
	for i := 0; i < len(base); i++ {
		if charSet[base[i]] {
			return false
		}
		charSet[base[i]] = true
	}
	return true
}

func isValidString(s string, base string) bool {
	for i := 0; i < len(s); i++ {
		isValid := false
		for j := 0; j < len(base); j++ {
			if s[i] == base[j] {
				isValid = true
				break
			}
		}
		if !isValid {
			return false
		}
	}
	return true
}
