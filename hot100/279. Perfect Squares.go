package hot100

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = i
	}
	for i := 2; i <= 100; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}
	return dp[n]
}
