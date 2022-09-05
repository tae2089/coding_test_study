package coding

func RomanToInt(s string) int {
    mapping := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var sum int = 0
	var prev rune = 0

	for _, v := range s{
		sum += mapping[v]
		if mapping[prev] < mapping[v] {
			sum -= mapping[prev] * 2
		}
		prev = v
	}
	return sum
}


