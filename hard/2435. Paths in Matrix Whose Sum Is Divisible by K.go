package hard

func numberOfPaths(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k)
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += grid[0][i]
		dp[i][sum%k] = 1
	}
	mod := int(1e9 + 7)
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j >= 1 {
				for l := 0; l < k; l++ {
					dp[j][l] += dp[j-1][l]
				}
			}
			curr := make([]int, k)
			for l := 0; l < k; l++ {
				if dp[j][l] > 0 {
					curr[(l+grid[i][j])%k] = dp[j][l] % mod
				}
			}
			copy(dp[j], curr)
		}
	}
	return dp[n-1][0]
}
