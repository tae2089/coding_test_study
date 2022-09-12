package coding

import "sort"

func FindIntArrayAboucDivisor(arr []int, divisor int)[]int {
	result := make([]int,0)
	for _,v := range arr {
		if v % divisor == 0 {
			result = append(result,v)
		}
	}
	if len(result) == 0 {
		return []int{-1}
	}
	sort.Ints(result)
	return result
}