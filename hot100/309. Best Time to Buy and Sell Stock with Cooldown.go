package hot100

import "math"

func maxProfit309(prices []int) int {
	dp := [3]int{math.MinInt32, 0, 0}
	for i := 0; i < len(prices); i++ {
		dp[0], dp[1], dp[2] = max(dp[0], dp[2]-prices[i]), dp[0]+prices[i], max(dp[1], dp[2])
	}
	return max(dp[1], dp[2])
}
