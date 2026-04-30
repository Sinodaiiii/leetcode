package medium

func maxPathScore260430(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	next := 1
	if grid[0][0] != 0 {
		dp[0][next] = grid[0][0]
		next += 1
	} else {
		dp[0][0] = 0
	}
	lnext := next
	for i := 1; i < n && lnext <= k; i++ {
		if grid[0][i] != 0 {
			dp[i][lnext] = dp[i-1][lnext-1] + grid[0][i]
			lnext += 1
		} else {
			dp[i][lnext-1] = dp[i-1][lnext-1]
		}
	}
	// fmt.Println(dp)
	for i := 1; i < m; i++ {
		if grid[i][0] != 0 {
			if next <= k {
				dp[0][next] = dp[0][next-1] + grid[i][0]
				dp[0][next-1] = -1
				next += 1
			} else {
				dp[0][k] = -1
			}
		}
		for j := 1; j < n; j++ {
			tmp := make([]int, k+1)
			for t := range tmp {
				tmp[t] = -1
			}
			for t := k; t >= 0; t-- {
				if grid[i][j] == 0 {
					tmp[t] = max(dp[j-1][t], dp[j][t])
				} else {
					if t == 0 {
						continue
					}
					pre := max(dp[j-1][t-1], dp[j][t-1])
					if pre >= 0 {
						tmp[t] = pre + grid[i][j]
					}
				}
			}
			dp[j] = tmp
		}
		// fmt.Println(dp)
	}
	ans := -1
	for i := 0; i <= k; i++ {
		if dp[n-1][i] >= 0 {
			ans = max(ans, dp[n-1][i])
		}
	}
	return ans
}
