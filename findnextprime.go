package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	prime := nb
	for {
		if IsPrime2(prime) {
			return prime
		}
		prime++
	}
}

func IsPrime2(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb%2 == 0 {
		return false
	}
	for i := 3; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
