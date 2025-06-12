package hot100

func climbStairs(n int) int {
	dp := make([]int, 2)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i%2] = dp[0] + dp[1]
	}
	return dp[n%2]
}
