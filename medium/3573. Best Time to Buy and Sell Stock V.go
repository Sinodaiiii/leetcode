package medium

import "math"

func maximumProfit(prices []int, k int) int64 {
	n := len(prices)
	dp := make([][][4]int, n)
	dp[0] = make([][4]int, k)
	dp[0][0] = [4]int{-prices[0], math.MinInt32, prices[0], math.MinInt32}
	for i := 1; i < k; i++ {
		dp[0][i] = [4]int{math.MinInt32, math.MinInt32, math.MinInt32, math.MinInt32}
	}
	for i := 1; i < n; i++ {
		dp[i] = make([][4]int, k)
		dp[i][0] = [4]int{max(-prices[i], dp[i-1][0][0]), max(dp[i-1][0][0]+prices[i], dp[i-1][0][1]), max(prices[i], dp[i-1][0][2]), max(dp[i-1][0][2]-prices[i], dp[i-1][0][3])}
	}
	// fmt.Println(dp[0])
	for i := 1; i < n; i++ {
		for j := 1; j < k; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], max(dp[i-1][j-1][1], dp[i-1][j-1][3])-prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j][0]+prices[i])
			dp[i][j][2] = max(dp[i-1][j][2], max(dp[i-1][j-1][1], dp[i-1][j-1][3])+prices[i])
			dp[i][j][3] = max(dp[i-1][j][3], dp[i-1][j][2]-prices[i])
		}
		// fmt.Println(dp[i])
	}
	ans := 0
	for i := 0; i < k; i++ {
		ans = max(ans, dp[n-1][i][1], dp[n-1][i][3])
	}
	return int64(ans)
}
