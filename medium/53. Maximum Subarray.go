package medium

func maxSubArray(nums []int) int {
	n := len(nums)
	dp, ans := nums[0], nums[0]
	for i := 1; i < n; i++ {
		dp = max(dp+nums[i], nums[i])
		ans = max(ans, dp)
	}
	return ans
}
