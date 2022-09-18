package coding


func Tonurment(n, a, b int) int{
	var answer int
	for{
		a = a / 2 + a % 2
		b = b / 2 + b % 2
		answer++
		if a == b{
			break
		}
	}
	return answer
}
