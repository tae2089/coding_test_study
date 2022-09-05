package coding

import "strconv"

func IsPalindrome(x int) bool {
	var res string
	var toString = strconv.Itoa(x)

	for i := len(toString) - 1; i >= 0; i-- {
		res = res + string(toString[i])
	}
	
	return res == toString
}
