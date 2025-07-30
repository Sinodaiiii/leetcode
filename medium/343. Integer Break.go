package medium

func integerBreak(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(max(dp[i], dp[j]*(i-j)), j*(i-j))
		}
	}
	return dp[n]
}
