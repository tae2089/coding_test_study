package coding

func FindMbti(survey []string, choices []int) string {
	mbtiType := map[string]int{
		"R":0,
		"T":0,
		"C":0, 
		"F":0,
		"J":0,
		"M":0,
		"A":0,
		"N":0,
	}

	for idx ,question := range survey{
			switch choices[idx]{
				case 1:
					mbtiType[string(question[0])] +=3
					
				case 2:
					mbtiType[string(question[0])] +=2
					
				case 3:
					mbtiType[string(question[0])] +=1
					
				case 5:
					mbtiType[string(question[1])] +=1
					
				case 6:
					mbtiType[string(question[1])] +=2
					
				case 7:
					mbtiType[string(question[1])] +=3
			}
	}
	firstChar := getMbti("R","T",mbtiType)
	secondChar := getMbti("C","F",mbtiType)
	thirdChar := getMbti("J","M",mbtiType)
	fourthChar := getMbti("A","N",mbtiType)
    return firstChar+secondChar+thirdChar+fourthChar
}

func getMbti(type1,type2 string, dict map[string]int) string{
	value1 := dict[type1]
	value2 := dict[type2]
	if value1 > value2 {
		return type1
	}else if value1 < value2 {
		return type2
	}else{
		return type1
	}
}