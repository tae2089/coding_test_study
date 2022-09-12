package coding

func IsNumber(s string) bool {
	if len(s) != 4 && len(s) !=6{
		return false
	}

	runes := []rune(s)
	for i:=0; i < len(runes); i++ {
		r := runes[i]
		if r > 96 && r < 123 {
			return false
		}else if r >65 && r<91{
			return false
		}
	}
	return true
}