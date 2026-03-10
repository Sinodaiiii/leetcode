package hard

func numberOfStableArrays260310(zero int, one int, limit int) int {
	dp := make([][][2]int, zero+1)
	for i := 0; i <= zero; i++ {
		dp[i] = make([][2]int, one+1)
	}
	for i := 1; i <= limit && i <= zero; i++ {
		dp[i][0][0] = 1
	}
	for i := 1; i <= limit && i <= one; i++ {
		dp[0][i][1] = 1
	}
	mod := 1000000007
	for i := 1; i <= zero; i++ {
		for j := 1; j <= one; j++ {
			zeroLimit := 0
			if i-limit-1 >= 0 {
				zeroLimit = dp[i-limit-1][j][1]
			}
			dp[i][j][0] = ((dp[i-1][j][0]+dp[i-1][j][1])%mod + 2*mod - zeroLimit) % mod

			oneLimit := 0
			if j-limit-1 >= 0 {
				oneLimit = dp[i][j-limit-1][0]
			}
			dp[i][j][1] = ((dp[i][j-1][0] + dp[i][j-1][1]) + 2*mod - oneLimit) % mod
		}
	}
	return (dp[zero][one][0] + dp[zero][one][1]) % mod
}
