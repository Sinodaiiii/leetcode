package hard

import "math"

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][]int, 4)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = math.MinInt32
		}
	}
	ans := 0
	symbol := -1
	dp[0][0] = symbol * prices[0]
	for j := 1; j < n; j++ {
		dp[0][j] = max(dp[0][j-1], symbol*prices[j])
	}
	for i := 1; i < 4; i++ {
		symbol *= -1
		for j := i; j < n; j++ {
			dp[i][j] = max(dp[i-1][j-1]+symbol*prices[j], dp[i][j-1])
			ans = max(dp[i][j], ans)
		}
	}
	return ans
}

func maxProfit260122(prices []int) int {
	dp := make([][2]int, 2)
	dp[0][0] = -prices[0]
	dp[0][1], dp[1][0], dp[1][1] = math.MinInt32, math.MinInt32, math.MinInt32
	for i := 1; i < len(prices); i++ {
		dp[1][1] = max(dp[1][1], dp[1][0]+prices[i])
		dp[1][0] = max(dp[1][0], dp[0][1]-prices[i])
		dp[0][1] = max(dp[0][1], dp[0][0]+prices[i])
		dp[0][0] = max(dp[0][0], -prices[i])
	}
	return max(dp[0][1], dp[1][1], 0)
}
