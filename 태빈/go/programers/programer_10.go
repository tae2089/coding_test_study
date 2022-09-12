package coding

func FindMinValue(arr []int)[]int{
	minValue := findMin(arr)
	result := make([]int,0)

	for _,value := range arr{
		if value != minValue{
			result = append(result, value)
		}
	}
	if len(result) == 0 {
		result = append(result, -1)
	}
	return result
}

func findMin(arr []int) int {
	minValue := 9999999

	for _, v := range arr {
		if v < minValue {
			minValue = v
		}
	}
	return minValue
}