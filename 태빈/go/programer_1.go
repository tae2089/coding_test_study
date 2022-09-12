package coding

func FindNumber(n int) int{
	var result int
	for i := 2; i < n; i++ {
		if(n % i == 1){
			result = i
			break
		}
	}
	return result
}