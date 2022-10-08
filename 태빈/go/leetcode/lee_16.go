package leetcode

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var answer int = 0
	var minDiff int = math.MaxInt32
	length := len(nums)

	for i := 0; i < length; i++ {

		l, h := i+1, length-1
		for l < h {
			s := nums[i] + nums[l] + nums[h]
			if abs(target-s) < minDiff {
				minDiff = abs(target - s)
				answer = s
			}
			if s > target {
				h--
			} else {
				l++
			}
		}
	}
	return answer
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
