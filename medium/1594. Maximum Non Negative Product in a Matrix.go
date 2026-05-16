package medium

func maxProductPath260516(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][2]int, n)
	dp[0] = [2]int{grid[0][0], grid[0][0]}
	mod := 1000000007
	for i := 1; i < n; i++ {
		minNum := dp[i-1][0] * grid[0][i]
		maxNum := dp[i-1][1] * grid[0][i]
		dp[i] = [2]int{min(minNum, maxNum), max(minNum, maxNum)}
	}
	for i := 1; i < m; i++ {
		minNum := dp[0][0] * grid[i][0]
		maxNum := dp[0][1] * grid[i][0]
		dp[0] = [2]int{min(minNum, maxNum), max(minNum, maxNum)}
		for j := 1; j < n; j++ {
			min1, max1 := dp[j][0]*grid[i][j], dp[j][1]*grid[i][j]
			min2, max2 := dp[j-1][0]*grid[i][j], dp[j-1][1]*grid[i][j]
			dp[j] = [2]int{min(min1, min2, max1, max2), max(min1, min2, max1, max2)}
		}
	}
	return max(dp[n-1][0]%mod, dp[n-1][1]%mod, -1)
}
