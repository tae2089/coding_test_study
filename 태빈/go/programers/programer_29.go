package coding

import "sort"

func RandomSorting(strings []string , n int)[] string{
	sort.Slice(strings, func(i, j int) bool {
		if strings[i][n] == strings[j][n] {
			return strings[i] < strings[j]
		}
		return strings[i][n] < strings[j][n]
	})
	return strings
}