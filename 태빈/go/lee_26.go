package coding

func RemoveDuplicates(nums []int) int {

	if len(nums) == 1 {
		return 1
	}

	idx := 1

	for i:=1; i<len(nums); i++ {
		if nums[idx -1] !=  nums[i] {
			nums[idx] = nums[i]
			idx++
		}
	}
	
	return idx
}