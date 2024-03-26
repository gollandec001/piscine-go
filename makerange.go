package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	length := max - min
	if length < 0 {
		length = 0
	}
	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[i] = min + i
	}
	return result
}
