package medium

import "math"

func maxSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	dp := make([]int, k)
	for i := 0; i < k; i++ {
		dp[i] = math.MinInt32 * 10
	}
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	ans := sum
	dp[0] = sum
	index := 0
	for i := k; i < n; i++ {
		index = (index + 1) % k
		sum = sum + nums[i] - nums[i-k]
		dp[index] = max(dp[index]+sum, sum)
		ans = max(ans, dp[index])
	}
	return int64(ans)
}
