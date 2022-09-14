package coding

func AddDivisor(left,right int) int {
	var result int
	for i:=left; i<right+1; i++ {
		length := countDivisors(i)
		if length % 2 == 0{
			result += i
		}else{
			result -=i
		}
	}
	return result
}

func countDivisors(num int) int{
	var count int
	for i := 1; i <= num; i++ {
		if num % i == 0{
			count++
		}
	}
	return count
}