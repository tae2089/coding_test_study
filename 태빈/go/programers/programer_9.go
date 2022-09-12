package coding

func JumpToX(x,n int) []int64 {
	var result []int64
	for i:= 1; i<=n;i++{
		value := i*x
		result = append(result, int64(value))
	}
	return result
}