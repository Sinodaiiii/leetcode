package hot100

func rob(nums []int) int {
	dp := [2]int{0, 0}
	for i := 0; i < len(nums); i++ {
		dp = [2]int{max(dp[0], dp[1]), dp[0] + nums[i]}
	}
	return max(dp[0], dp[1])
}
