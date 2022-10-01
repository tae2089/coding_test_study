package leetcode

func numDecodings(s string) int {
	dp := make([]int, len(s))
	if s[0] == '0' {
		return 0
	}
	dp[0] = 1
	for i := 1; i < len(s); i++ {
		curOneSum := int(s[i] - '0')
		curTenSum := int(s[i-1]-'0')*10 + curOneSum
		if curOneSum != 0 {
			dp[i] += dp[i-1]
		}
		if curTenSum >= 10 && curTenSum <= 26 {
			if i >= 2 {
				dp[i] += dp[i-2]
			} else {
				dp[i] += 1
			}
		}
	}
	return dp[len(s)-1]
}
