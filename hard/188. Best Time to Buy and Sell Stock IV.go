package hard

import "math"

func maxProfit188(k int, prices []int) int {
	n := len(prices)
	dp := make([][2]int, k)
	dp[0] = [2]int{-prices[0], math.MinInt32}
	for i := 1; i < k; i++ {
		dp[i] = [2]int{math.MinInt32, math.MinInt32}
	}
	for i := 1; i < n; i++ {
		for j := k - 1; j > 0; j-- {
			dp[j][0] = max(dp[j][0], dp[j-1][1]-prices[i])
			dp[j][1] = max(dp[j][1], dp[j][0]+prices[i])
		}
		dp[0] = [2]int{max(dp[0][0], -prices[i]), max(dp[0][1], dp[0][0]+prices[i])}
	}
	ans := 0
	for i := 0; i < k; i++ {
		ans = max(ans, dp[i][1])
	}
	return ans
}
