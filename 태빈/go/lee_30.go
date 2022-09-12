package coding

//sliding window algorithm
func LengthOfLongestSubstring(s string) int {
	strLength := len(s)

	if strLength == 0 {
		return 0
	}
	// left point, right point init
	lp := 0
	rp := 0

	//max lenth variable init
	maxLength := -1

	//window init
	window := make(map[byte]int, strLength)

	for rp < strLength{
		b := s[rp]
		rp++
		window[b]++

		for window[b] == 2 {
			c := s[lp]
			window[c]--
			lp++
		}
				// update maxLength
		if (rp - lp) > maxLength {
			maxLength = rp - lp
		}
	}
	return maxLength
}