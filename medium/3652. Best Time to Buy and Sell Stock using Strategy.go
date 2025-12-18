package medium

func maxProfit3652(prices []int, strategy []int, k int) int64 {
	n := len(prices)
	profit := make([]int, n)
	curr := 0
	for i := 0; i <= n-k; i++ {
		profit[i] = curr
		curr += prices[i] * strategy[i]
	}
	for i := n - k + 1; i < n; i++ {
		curr += prices[i] * strategy[i]
	}
	ans := curr
	// fmt.Println(profit)
	curr = 0
	for i := n - 1; i >= k-1; i-- {
		profit[i-k+1] += curr
		curr += prices[i] * strategy[i]
	}
	// fmt.Println(profit)
	curr = 0
	for i := k / 2; i < k; i++ {
		curr += prices[i]
	}
	ans = max(ans, profit[0]+curr)
	for i := 1; i <= n-k; i++ {
		// fmt.Println(i, curr)
		curr = curr - prices[i+k/2-1] + prices[i+k-1]
		// fmt.Println(curr)
		ans = max(ans, curr+profit[i])
	}
	// fmt.Println(profit)
	return int64(ans)
}
