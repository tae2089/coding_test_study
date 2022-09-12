package coding

func CalculateScarceMoney(price int, money int, count int) int64 {

	var result int
	for i := 1; i< count+1; i++{
		result += price * i
	}
	if money  > result {
		return 0
	}
    return int64(result - money)
}