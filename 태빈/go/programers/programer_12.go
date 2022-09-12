package coding

import (
	"sort"
	"strconv"
)
func ReversInt(n int64)int64{
	strN := strconv.FormatInt(n, 10)
	runes := []rune(strN)
	sort.Slice(runes,func(i, j int) bool {
		return runes[j] < runes[i]
	})
	strN = string(runes)
	result, _:= strconv.ParseInt(strN, 0, 64)
	return result
}