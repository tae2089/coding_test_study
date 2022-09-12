package coding

import "strconv"

func StrToInt(s string) int {
	num,_ := strconv.Atoi(s)
	return num
}
