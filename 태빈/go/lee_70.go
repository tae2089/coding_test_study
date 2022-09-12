package coding

func ClimbStairs(n int) int {

	if n < 3 {
		return n
	}
	first := 1
	second := 2

	for i:=3; i<=n;i++{
		first, second = second, first+second
	}
	
	return second
}