package medium

func maxSubarraySumCircular(nums []int) int {
	dp := nums[0]
	ans := dp
	n := len(nums)
	for i := 1; i < n*2; i++ {
		j := i % n
		curr := dp + nums[j]
		if i >= n {
			curr -= nums[]
		}
		dp = max(dp+nums[j], nums[j])
		ans = max(ans, dp)
	}
	return ans
}
