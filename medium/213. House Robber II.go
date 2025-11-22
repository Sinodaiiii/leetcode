package medium

func rob213(nums []int) int {
	n := len(nums)
	dp := [2]int{nums[0], 0}
	for i := 1; i < n-1; i++ {
		dp = [2]int{dp[1] + nums[i], max(dp[1], dp[0])}
	}
	ans := max(dp[0], dp[1])
	dp = [2]int{0, nums[n-1]}
	for i := 1; i < n-2; i++ {
		dp = [2]int{dp[1] + nums[i], max(dp[1], dp[0])}
	}
	ans = max(ans, dp[0], dp[1])
	return ans
}
