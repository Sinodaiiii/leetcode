package medium

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([]int, 3)
	buyi := 0
	dp[0], dp[1], dp[2] = 0, prices[0], 0
	if n >= 2 {
		dp[1] = max(dp[0], dp[1]+n, dp[2])
	}
	for i := 1; i < n; i++ {
		dp[]
	}
	return max(dp[0], dp[1], dp[2])
}
