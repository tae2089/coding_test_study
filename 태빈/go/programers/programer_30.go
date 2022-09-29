package coding

func RentCloth(n int, lost []int, reserve []int) int {
	student := [32]int{0,}
	var result int
	for _,v := range reserve{
		student[v] +=1
	}

	for _,v := range lost{
		student[v] -=1
	}

	for i:= 1; i<=n; i++{
		if student[i] == -1{
			if student[i - 1] == 1{
				student[i - 1] = 0
				student[i] = 0
			}else if student[i +1 ] == 1{
				student[i + 1] = 0
				student[i] = 0
			}
		}
	}

	for i:=1;i<=n;i++{
		if student[i] != -1{
			result++
		}
	}

    return  result
}