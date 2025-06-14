package hot100

func maxProfit(prices []int) int {
	m, n := prices[0], len(prices)
	ans := 0
	for i := 1; i < n; i++ {
		if prices[i] < m {
			m = prices[i]
		} else {
			ans = max(prices[i]-m, ans)
		}
	}
	return ans
}
