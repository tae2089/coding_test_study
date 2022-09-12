package coding

import "math"

func IsSqrt(n int64) int64{
	
	sqrt := int64(math.Sqrt(float64(n)))
	if sqrt*sqrt == n {
		sqrt++
		return sqrt * sqrt
	}

	return -1
}