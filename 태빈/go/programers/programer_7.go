package coding

func Collas(num int) int{
	var count int

	if num == 1 {
		return 0
	}

	for i := 0; i < 501; i++ {
		if num == 1 {
			return count
		}
		if num % 2 == 0{
			num = num / 2
		}else{
			num = num * 3 +1
		}
		count++
	}

	return -1
}