package coding

import "strconv"

func IsHashard(n int) bool{
	strN := strconv.Itoa(n)
	runes := []rune(strN)
	hashardNum := 0

	for i:=0; i < len(runes); i++ {
		r:=runes[i]
		hashardNum += int(r - '0')
	}

	if n % hashardNum == 0{
		return true
	}else{
		return false
	}
}