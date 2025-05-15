package problems

func numDistinct(s string, t string) int {
	dp := make([]int, len(s))
	for i := 0; i < len(t); i++ {
		num := 0
		for j := 0; j < len(s); j++ {
			if s[j] == t[i] {
				if i == 0 {
					dp[j] = 1
				} else {
					tmp := dp[j]
					dp[j] = num
					num += tmp
				}
			} else if dp[j] > 0 {
				num += dp[j]
				dp[j] = 0
			}
		}
	}
	ans := 0
	for i := 0; i < len(dp); i++ {
		ans += dp[i]
	}
	return ans
}
