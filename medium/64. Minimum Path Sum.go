package medium

import "math"

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < n; i++ {
		grid[0][i] += grid[0][i-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] = min(grid[i-1][j]+grid[i][j], grid[i][j-1]+grid[i][j])
		}
	}
	return grid[m-1][n-1]
}

func minPathSum2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt32
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				dp[j] += grid[i][j]
			} else {
				dp[j] = min(dp[j-1], dp[j]) + grid[i][j]
			}
		}
	}
	return dp[n-1]
}
