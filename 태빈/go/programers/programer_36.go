package coding


func Fibonaci(n  int) int {
	
	answer := []int{0,1}
	
	for i := 2; i <=n ; i++ {
        result := (answer[i - 1] + answer[i - 2]) % 1234567
		answer = append(answer,result)
	}
	
    return answer[len(answer)-1] 
}