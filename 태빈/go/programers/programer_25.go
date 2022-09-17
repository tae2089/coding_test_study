package coding

//유클리드 호제법
func GCDandLCM(n,m int)[]int{
	a := GCD(n,m)
	b := lcm(n,m,a)
	return []int{a,b}
}

func GCD(a, b int) int {
	if a < b {
		a, b = b, a
	}

	gcd := 0
	for {
		r := a % b
		if r == 0 {
			gcd = b
			break
		}
		a = b
		b = r
	}

	return gcd
}

func lcm(a, b,gcd int) int {
	result := (a*b) / gcd
	return result
}