package medium

import "math"

func maxProfit(prices []int) int {
	dp := [3]int{math.MinInt32, 0, 0}
	for i := 0; i < len(prices); i++ {
		dp[0], dp[1], dp[2] = max(dp[0], dp[2]-prices[i]), dp[0]+prices[i], max(dp[1], dp[2])
	}
	return max(dp[1], dp[2])
}

func maxProfit260220(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0] = [2]int{-prices[0], 0}
	if n >= 2 {
		dp[1] = [2]int{max(dp[0][0], -prices[1]), max(0, prices[1]-prices[0])}
	}
	// fmt.Println(dp)
	for i := 2; i < n; i++ {
		dp[i] = [2]int{max(dp[i-1][0], dp[i-2][1]-prices[i]), max(dp[i-1][0]+prices[i], dp[i-1][1])}
		// fmt.Println(dp[i])
	}
	return max(dp[n-1][0], dp[n-1][1])
}
