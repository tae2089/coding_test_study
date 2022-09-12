package coding

func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] = digits[i] + 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	result := []int{}
	if digits[0] == 0 {
		result = append([]int{1}, digits...)
	}

	return result
}