package coding

func AddArray(arr1,arr2 [][]int) [][]int{
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[i]);j++{
			arr1[i][j] = arr1[i][j]+arr2[i][j]		
		}
	}
	return arr1
}
