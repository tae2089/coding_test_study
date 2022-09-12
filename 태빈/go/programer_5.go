package coding

import (
	"strconv"
)


func ReversIntarray(n int64)[]int{
	var result []int
	strN := strconv.FormatInt(n, 10)
	runes := []rune(strN)

	for i := len(runes)-1; i >=0; i-- {
		result = append(result, int(runes[i] - '0'))
	}
	return result
}