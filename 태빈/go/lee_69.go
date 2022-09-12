package coding

func Mysqrt(x int)int{
	var output = 0

	for output * output < x{
		output++
	}

	if output * output > x{
		output--
	}
	
	return output
}