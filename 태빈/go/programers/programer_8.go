package coding

import "fmt"

func FindKim(seoul []string) string {
	var findIndex int
	for idx,value := range seoul {
		if value == "Kim"{
			findIndex = idx
			break
		}
	}
    return fmt.Sprintf("김서방은 %d에 있다",findIndex)
}