package coding

import "strconv"

func AddStringNumbers(n int )int {

	var result int
	strN := strconv.Itoa(n)

	for _,v := range strN {
		num ,_:= strconv.Atoi(string(v))
		result += num
	}

	return result
}