package coding

import "sort"

func SortSentance(s string) string {
	
	runes := []rune(s)
	smallArray := make([]rune,0)
	bigArray := make([]rune,0)
	
	for i:=0; i < len(runes); i++{
		value := runes[i]
		if value >96 && value<123{
			smallArray = append(smallArray,value)
		}else{
			bigArray = append(bigArray,value)
		}
	}

	sort.Slice(smallArray, func(i, j int) bool {
		return smallArray[i] > smallArray[j]
	})
	sort.Slice(bigArray, func(i, j int) bool{
		return bigArray[i] > bigArray[j]
	})
    return string(smallArray)+string(bigArray)
}