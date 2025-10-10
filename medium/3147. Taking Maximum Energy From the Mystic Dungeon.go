package medium

import "math"

func maximumEnergy(energy []int, k int) int {
	dp := make([]int, k)
	for i := 0; i < k; i++ {
		dp[i] = math.MinInt32
	}
	n := len(energy)
	for i := 0; i < n; i++ {
		dp[i%k] = max(dp[i%k]+energy[i], energy[i])
	}
	ans := math.MinInt32
	for i := 0; i < k; i++ {
		ans = max(ans, dp[i])
	}
	return ans
}
