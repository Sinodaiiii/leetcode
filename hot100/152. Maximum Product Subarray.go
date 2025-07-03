package hot100

func maxProduct(nums []int) int {
	ans := nums[0]
	n := len(nums)
	dp := [2]int{1, 1}
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			ans = max(ans, 0)
			dp = [2]int{1, 1}
		} else if nums[i] > 0 {
			dp[0] *= nums[i]
			dp[1] = min(dp[1]*nums[i], 1)
			ans = max(dp[0], ans)
		} else {
			if dp[1] < 0 {
				ans = max(dp[1]*nums[i], ans)
			}
			dp[0], dp[1] = max(dp[1]*nums[i], 1), dp[0]*nums[i]
		}
	}
	return ans
}
