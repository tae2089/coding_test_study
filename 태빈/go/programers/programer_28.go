package coding

func GetDate(a,b int) string{
	day := []string{"FRI", "SAT", "SUN", "MON", "TUE", "WED", "THU"}
	month := []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	date := 0
	for i:=0 ; i<a-1 ;i ++{
		date += month[i]
	}
	date += b-1;
	return day[date%7]
}