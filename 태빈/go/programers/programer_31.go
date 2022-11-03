package coding

func Lotto(lottos []int, win_nums []int) []int {
	// key - 맞은 갯수 , value - 순위
	lottoRank := map[int]int{
		6:1,
		5:2,
		4:3,
		3:4,
		2:5,
		1:6,
		0:6,
	}
	lottoDic := make(map[int]int)
	zeroCnt := 0
	result := 0
	for _, num := range lottos {
		if(num == 0){
			zeroCnt++
			continue
		}
		lottoDic[num] = 1
	}

	for _, num := range win_nums {
		lottoDic[num] += 1
		if(lottoDic[num] == 2){
			result+=1
		}
	}
	
	min := lottoRank[result]
	max := lottoRank[result+zeroCnt]
    return []int{max,min}
}