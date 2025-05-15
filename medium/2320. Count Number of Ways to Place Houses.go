package medium

func countHousePlacements(n int) int {
	if n == 1 {
		return 4
	}
	mod := int(1e9 + 7)
	dp := []int{1, 1, 1, 1}
	for i := 2; i <= n; i++ {
		dp[0], dp[1], dp[2], dp[3] = (dp[0]+dp[1]+dp[2]+dp[3])%mod, (dp[0]+dp[2])%mod, (dp[0]+dp[1])%mod, dp[0]%mod
	}
	return (dp[0] + dp[1] + dp[2] + dp[3]) % mod
}
